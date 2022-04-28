/*
 * Copyright 2021 SkyAPM
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */


#include "sky_core_span.h"
#include "php.h"
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>
#include "sky_util_php.h"

sky_core_span_t *sky_core_span_new(sky_core_span_type type, sky_core_span_layer layer, int componentId) {
    sky_core_span_t *span = (sky_core_span_t *) emalloc(sizeof(sky_core_span_t));
    span->refs_total = 4;
    span->refs_size = 0;
    span->tags_total = 4;
    span->tags_size = 0;
    span->logs_total = 4;
    span->logs_size = 0;

    struct timeval tv;
    gettimeofday(&tv, NULL);
    span->startTime = tv.tv_sec * 1000 + tv.tv_usec / 1000;
    span->refs = (sky_core_segment_ref_t **) emalloc(span->refs_total);
    span->peer = (char *) emalloc(512);
    bzero(span->peer, 512);
    span->spanType = type;
    span->spanLayer = layer;
    span->componentId = componentId;

    span->isError = false;
    span->tags = (sky_core_tag_t **) emalloc(span->tags_total);
    span->logs = (sky_core_log_t **) emalloc(span->logs_total);
    span->skipAnalysis = false;
    return span;
}

void sky_core_span_set_end_time(sky_core_span_t *span) {
    struct timeval tv;
    gettimeofday(&tv, NULL);
    span->endTime = tv.tv_sec * 1000 + tv.tv_usec / 1000;
}

void sky_core_span_add_refs(sky_core_span_t *span, sky_core_segment_ref_t *ref) {
    span->refs[span->refs_size] = ref;
}

void sky_core_span_set_operation_name(sky_core_span_t *span, char *name) {
    span->operationName = (char *) emalloc(sizeof(char) * strlen(name));

    char *index = strrchr(name, '?');
    if (index != NULL) {
        strncpy(span->operationName, name, (int) (index - name));
    } else {
        strcpy(span->operationName, name);
    }
}

void sky_core_span_set_peer(sky_core_span_t *span, char *peer) {
    if (peer == NULL) {
        span->peer = "";
        return;
    }
    span->peer = peer;
}

void sky_core_span_set_error(sky_core_span_t *span, bool isError) {
    span->isError = isError;
}

void sky_core_span_add_tag(sky_core_span_t *span, sky_core_tag_t *tag) {
    if (span->tags_size == span->tags_total - 1) {
        sky_core_tag_t **more = (sky_core_tag_t **) erealloc(span->tags, span->tags_total * 2);
        if (more != NULL) {
            span->tags_total *= 2;
            span->tags = more;
        } else {
            return;
        }
    }

    span->tags[span->tags_size] = tag;
    span->tags_size++;
}

char *sky_core_span_to_json(sky_core_span_t *span) {
    char *json;
    char *temp = "{"
                 "\"spanId\":%d,"
                 "\"parentSpanId\":%d,"
                 "\"startTime\":%ld,"
                 "\"endTime\":%ld,"
                 "\"refs\":%s,"
                 "\"operationName\":\"%s\","
                 "\"peer\":\"%s\","
                 "\"spanType\":%d,"
                 "\"spanLayer\":%d,"
                 "\"componentId\":%d,"
                 "\"isError\":%s,"
                 "\"tags\":%s,"
                 "\"logs\":%s,"
                 "\"skipAnalysis\":%s"
                 "}";

    sky_util_smart_string tags = {0};
    sky_util_smart_string_appendl(&tags, "[", 1);
    for (int i = 0; i < span->tags_size; ++i) {
        sky_core_tag_t *tag = span->tags[i];
        char *tag_json = sky_core_tag_to_json(tag);
        sky_util_smart_string_appendl(&tags, tag_json, strlen(tag_json));
        if (i + 1 < span->tags_size) {
            sky_util_smart_string_appendl(&tags, ",", 1);
        }
    }
    sky_util_smart_string_appendl(&tags, "]", 1);
    sky_util_smart_string_0(&tags);

    sky_util_smart_string logs = {0};
    sky_util_smart_string_appendl(&logs, "[", 1);
    for (int i = 0; i < span->logs_size; ++i) {
        sky_core_log_t *log = span->logs[i];
        char *log_json = sky_core_log_to_json(log);
        sky_util_smart_string_appendl(&logs, log_json, strlen(log_json));
        if (i + 1 < span->logs_size) {
            sky_util_smart_string_appendl(&logs, ",", 1);
        }
    }
    sky_util_smart_string_appendl(&logs, "]", 1);
    sky_util_smart_string_0(&logs);


    asprintf(&json, temp,
             span->spanId,
             span->parentSpanId,
             span->startTime,
             span->endTime,
             "[]", // refs
             span->operationName,
             span->peer,
             span->spanType,
             span->spanLayer,
             span->componentId,
             span->isError ? "true" : "false",
             sky_util_smart_string_to_char(tags),
             sky_util_smart_string_to_char(logs),
             span->skipAnalysis ? "true" : "false"
    );
    return json;
}

//SkyCoreSpan::SkyCoreSpan() {
//    startTime = getUnixTimeStamp();
//    isError = false;
//    skipAnalysis = false;
//}
//
//SkyCoreSpan::~SkyCoreSpan() {
//    for (auto ref : refs) {
//        delete ref;
//    }
//    refs.clear();
//    refs.shrink_to_fit();
//
//    for (auto tag : tags) {
//        delete tag;
//    }
//    tags.clear();
//    tags.shrink_to_fit();
//
//    for (auto log : logs) {
//        delete log;
//    }
//    logs.clear();
//    logs.shrink_to_fit();
//}
//
//// get
//int SkyCoreSpan::getSpanId() const {
//    return spanId;
//}
//
//int SkyCoreSpan::getParentSpanId() const {
//    return parentSpanId;
//}
//
//long SkyCoreSpan::getStartTime() const {
//    return startTime;
//}
//
//long SkyCoreSpan::getEndTime() const {
//    return endTime;
//}
//
//const std::vector<SkyCoreSegmentReference*>& SkyCoreSpan::getRefs() {
//    return refs;
//}
//
//const std::string& SkyCoreSpan::getOperationName() {
//    return operationName;
//}
//
//const std::string& SkyCoreSpan::getPeer() {
//    return peer;
//}
//
//SkyCoreSpanType SkyCoreSpan::getSpanType() {
//    return spanType;
//}
//
//SkyCoreSpanLayer SkyCoreSpan::getSpanLayer() {
//    return spanLayer;
//}
//
//int SkyCoreSpan::getComponentId() const {
//    return componentId;
//}
//
//bool SkyCoreSpan::getIsError() const {
//    return isError;
//}
//
//const std::vector<SkyCoreTag*>& SkyCoreSpan::getTags() {
//    return tags;
//}
//
//const std::vector<SkyCoreLog*>& SkyCoreSpan::getLogs() {
//    return logs;
//}
//
//bool SkyCoreSpan::getSkipAnalysis() const {
//    return skipAnalysis;
//}
//
//// set
//void SkyCoreSpan::setEndTIme() {
//    endTime = getUnixTimeStamp();
//}
//
//void SkyCoreSpan::setOperationName(const std::string &name) {
//    operationName = name.substr(0, name.find('?'));
//}
//
//void SkyCoreSpan::setPeer(const std::string &p) {
//    peer = p;
//}
//
//void SkyCoreSpan::setSpanType(SkyCoreSpanType type) {
//    spanType = type;
//}
//
//void SkyCoreSpan::setSpanLayer(SkyCoreSpanLayer layer) {
//    spanLayer = layer;
//}
//
//void SkyCoreSpan::setComponentId(int id) {
//    componentId = id;
//}
//
//void SkyCoreSpan::setIsError(bool error) {
//    isError = error;
//}
//
//void SkyCoreSpan::setSpanId(int id) {
//    spanId = id;
//}
//
//void SkyCoreSpan::setParentSpanId(int id) {
//    parentSpanId = id;
//}
//
//void SkyCoreSpan::pushTag(SkyCoreTag *tag) {
//    tags.push_back(tag);
//}
//
//void SkyCoreSpan::pushLog(SkyCoreLog *log) {
//    logs.push_back(log);
//}
//
//void SkyCoreSpan::addTag(const std::string &key, const std::string &value) {
//    tags.push_back(sky_core_tag_new(key.c_str(), value.c_str()));
//}
//
//void SkyCoreSpan::addLog(const std::string &key, const std::string &value) {
//    logs.push_back(new SkyCoreLog(key, value));
//}
//
//void SkyCoreSpan::pushRefs(SkyCoreSegmentReference *ref) {
//    refs.push_back(ref);
//}
//
//std::string SkyCoreSpan::marshal() {
//    nlohmann::json j;
//    j["spanId"] = spanId;
//    j["parentSpanId"] = parentSpanId;
//    j["startTime"] = startTime;
//    j["endTime"] = endTime;
//    j["refs"] = "todo";
//    j["operationName"] = operationName;
//    j["peer"] = peer;
//    j["spanType"] = spanType;
//    j["spanLayer"] = spanLayer;
//    j["componentId"] = componentId;
//    j["isError"] = isError;
//    j["tags"] = "todo";
//    j["logs"] = "todo";
//    j["skipAnalysis"] = skipAnalysis;
//    return j;
//}
///
//  Generated code. Do not modify.
//  source: topic.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const TopicReq$json = const {
  '1': 'TopicReq',
  '2': const [
    const {'1': 'topic', '3': 1, '4': 3, '5': 9, '10': 'topic'},
  ],
};

const TopicResp$json = const {
  '1': 'TopicResp',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'status', '3': 2, '4': 1, '5': 14, '6': '.api.Status', '10': 'status'},
  ],
};

const TopicServiceBase$json = const {
  '1': 'Topic',
  '2': const [
    const {'1': 'subscribe', '2': '.api.TopicReq', '3': '.api.TopicResp'},
    const {'1': 'unsubscribe', '2': '.api.TopicReq', '3': '.api.TopicResp'},
  ],
};

const TopicServiceBase$messageJson = const {
  '.api.TopicReq': TopicReq$json,
  '.api.TopicResp': TopicResp$json,
};


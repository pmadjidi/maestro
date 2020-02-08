///
//  Generated code. Do not modify.
//  source: topics.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'register.pbjson.dart' as $3;

const Topic$json = const {
  '1': 'Topic',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'tag', '3': 2, '4': 1, '5': 9, '10': 'tag'},
    const {'1': 'status', '3': 3, '4': 1, '5': 14, '6': '.api.Status', '10': 'status'},
  ],
};

const TopicReq$json = const {
  '1': 'TopicReq',
  '2': const [
    const {'1': 'list', '3': 1, '4': 3, '5': 11, '6': '.api.Topic', '10': 'list'},
  ],
};

const TopicResp$json = const {
  '1': 'TopicResp',
  '2': const [
    const {'1': 'list', '3': 1, '4': 3, '5': 11, '6': '.api.Topic', '10': 'list'},
  ],
};

const SubscriptionsServiceBase$json = const {
  '1': 'Subscriptions',
  '2': const [
    const {'1': 'sub', '2': '.api.TopicReq', '3': '.api.TopicResp'},
    const {'1': 'unsub', '2': '.api.TopicReq', '3': '.api.TopicResp'},
    const {'1': 'list', '2': '.api.Empty', '3': '.api.TopicResp'},
  ],
};

const SubscriptionsServiceBase$messageJson = const {
  '.api.TopicReq': TopicReq$json,
  '.api.Topic': Topic$json,
  '.api.TopicResp': TopicResp$json,
  '.api.Empty': $3.Empty$json,
};


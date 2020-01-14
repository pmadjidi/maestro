///
//  Generated code. Do not modify.
//  source: message.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'google/protobuf/timestamp.pbjson.dart' as $3;

const MsgReq$json = const {
  '1': 'MsgReq',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'text', '3': 2, '4': 1, '5': 9, '10': 'text'},
    const {'1': 'pic', '3': 3, '4': 1, '5': 12, '10': 'pic'},
    const {'1': 'parentId', '3': 4, '4': 1, '5': 9, '10': 'parentId'},
    const {'1': 'topic', '3': 5, '4': 1, '5': 9, '10': 'topic'},
    const {'1': 'time_name', '3': 6, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'timeName'},
    const {'1': 'status', '3': 7, '4': 1, '5': 14, '6': '.api.Status', '10': 'status'},
  ],
};

const MessageServiceBase$json = const {
  '1': 'Message',
  '2': const [
    const {'1': 'Msg', '2': '.api.MsgReq', '3': '.api.MsgReq', '5': true, '6': true},
  ],
};

const MessageServiceBase$messageJson = const {
  '.api.MsgReq': MsgReq$json,
  '.google.protobuf.Timestamp': $3.Timestamp$json,
};


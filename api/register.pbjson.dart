///
//  Generated code. Do not modify.
//  source: register.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'google/protobuf/timestamp.pbjson.dart' as $3;
import 'status.pbjson.dart' as $0;

const RegisterReq$json = const {
  '1': 'RegisterReq',
  '2': const [
    const {'1': 'userName', '3': 1, '4': 1, '5': 9, '10': 'userName'},
    const {'1': 'passWord', '3': 2, '4': 1, '5': 12, '10': 'passWord'},
    const {'1': 'firstName', '3': 3, '4': 1, '5': 9, '10': 'firstName'},
    const {'1': 'lastName', '3': 4, '4': 1, '5': 9, '10': 'lastName'},
    const {'1': 'email', '3': 5, '4': 1, '5': 9, '10': 'email'},
    const {'1': 'phone', '3': 6, '4': 1, '5': 9, '10': 'phone'},
    const {'1': 'address', '3': 7, '4': 1, '5': 11, '6': '.api.RegisterReq.Address', '10': 'address'},
    const {'1': 'device', '3': 8, '4': 1, '5': 9, '10': 'device'},
    const {'1': 'time_name', '3': 9, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'timeName'},
    const {'1': 'AppName', '3': 10, '4': 1, '5': 9, '10': 'AppName'},
  ],
  '3': const [RegisterReq_Address$json],
};

const RegisterReq_Address$json = const {
  '1': 'Address',
  '2': const [
    const {'1': 'street', '3': 1, '4': 1, '5': 9, '10': 'street'},
    const {'1': 'city', '3': 2, '4': 1, '5': 9, '10': 'city'},
    const {'1': 'state', '3': 3, '4': 1, '5': 9, '10': 'state'},
    const {'1': 'zip', '3': 4, '4': 1, '5': 9, '10': 'zip'},
  ],
};

const RegisterServiceBase$json = const {
  '1': 'Register',
  '2': const [
    const {'1': 'Register', '2': '.api.RegisterReq', '3': '.api.Empty'},
  ],
};

const RegisterServiceBase$messageJson = const {
  '.api.RegisterReq': RegisterReq$json,
  '.api.RegisterReq.Address': RegisterReq_Address$json,
  '.google.protobuf.Timestamp': $3.Timestamp$json,
  '.api.Empty': $0.Empty$json,
};


///
//  Generated code. Do not modify.
//  source: login.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'register.pbjson.dart' as $3;

const LoginReq$json = const {
  '1': 'LoginReq',
  '2': const [
    const {'1': 'device', '3': 1, '4': 1, '5': 9, '10': 'device'},
  ],
};

const LoginServiceBase$json = const {
  '1': 'Login',
  '2': const [
    const {'1': 'Authenticate', '2': '.api.LoginReq', '3': '.api.Empty'},
  ],
};

const LoginServiceBase$messageJson = const {
  '.api.LoginReq': LoginReq$json,
  '.api.Empty': $3.Empty$json,
};


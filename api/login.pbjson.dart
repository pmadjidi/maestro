///
//  Generated code. Do not modify.
//  source: login.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const LoginReq$json = const {
  '1': 'LoginReq',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'userName', '3': 2, '4': 1, '5': 9, '10': 'userName'},
    const {'1': 'passWord', '3': 3, '4': 1, '5': 12, '10': 'passWord'},
    const {'1': 'device', '3': 4, '4': 1, '5': 9, '10': 'device'},
  ],
};

const LoginResp$json = const {
  '1': 'LoginResp',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'status', '3': 2, '4': 1, '5': 14, '6': '.api.Status', '10': 'status'},
  ],
};

const LoginServiceBase$json = const {
  '1': 'Login',
  '2': const [
    const {'1': 'Authenticate', '2': '.api.LoginReq', '3': '.api.LoginResp'},
  ],
};

const LoginServiceBase$messageJson = const {
  '.api.LoginReq': LoginReq$json,
  '.api.LoginResp': LoginResp$json,
};


///
//  Generated code. Do not modify.
//  source: register.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

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

const RegisterResp$json = const {
  '1': 'RegisterResp',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'status', '3': 2, '4': 1, '5': 14, '6': '.api.Status', '10': 'status'},
  ],
};

const RegisterServiceBase$json = const {
  '1': 'Register',
  '2': const [
    const {'1': 'Register', '2': '.api.RegisterReq', '3': '.api.RegisterResp'},
  ],
};

const RegisterServiceBase$messageJson = const {
  '.api.RegisterReq': RegisterReq$json,
  '.api.RegisterReq.Address': RegisterReq_Address$json,
  '.api.RegisterResp': RegisterResp$json,
};


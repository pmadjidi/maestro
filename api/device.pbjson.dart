///
//  Generated code. Do not modify.
//  source: device.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const DeviceReq$json = const {
  '1': 'DeviceReq',
  '2': const [
    const {'1': 'deviceId', '3': 1, '4': 1, '5': 9, '10': 'deviceId'},
    const {'1': 'deviceKey', '3': 2, '4': 1, '5': 9, '10': 'deviceKey'},
  ],
};

const DeviceResp$json = const {
  '1': 'DeviceResp',
  '2': const [
    const {'1': 'status', '3': 1, '4': 1, '5': 14, '6': '.api.Status', '10': 'status'},
  ],
};

const DeviceServiceBase$json = const {
  '1': 'Device',
  '2': const [
    const {'1': 'Register', '2': '.api.DeviceReq', '3': '.api.DeviceResp'},
  ],
};

const DeviceServiceBase$messageJson = const {
  '.api.DeviceReq': DeviceReq$json,
  '.api.DeviceResp': DeviceResp$json,
};


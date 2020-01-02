///
//  Generated code. Do not modify.
//  source: status.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Status extends $pb.ProtobufEnum {
  static const Status SUCCESS = Status._(0, 'SUCCESS');
  static const Status FAIL = Status._(1, 'FAIL');
  static const Status BLOCKED = Status._(2, 'BLOCKED');
  static const Status DELETED = Status._(3, 'DELETED');
  static const Status CREATED = Status._(4, 'CREATED');
  static const Status TIMEOUT = Status._(5, 'TIMEOUT');
  static const Status ERROR = Status._(6, 'ERROR');
  static const Status NOTFOUND = Status._(7, 'NOTFOUND');
  static const Status INVALID_PASSWORD = Status._(8, 'INVALID_PASSWORD');
  static const Status INVALID_FIRSTNAME = Status._(9, 'INVALID_FIRSTNAME');
  static const Status INVALID_LASTNAME = Status._(10, 'INVALID_LASTNAME');
  static const Status INVALID_EMAIL = Status._(11, 'INVALID_EMAIL');
  static const Status INVALID_PHONE = Status._(12, 'INVALID_PHONE');
  static const Status INVALID_ADRESS = Status._(13, 'INVALID_ADRESS');
  static const Status INVALID_DEVICE = Status._(14, 'INVALID_DEVICE');
  static const Status INVALID_USERNAME = Status._(15, 'INVALID_USERNAME');
  static const Status VALIDATED = Status._(16, 'VALIDATED');
  static const Status EXITSTS = Status._(17, 'EXITSTS');
  static const Status MAXIMUN_NUMBER_OF_USERS_REACHED = Status._(18, 'MAXIMUN_NUMBER_OF_USERS_REACHED');

  static const $core.List<Status> values = <Status> [
    SUCCESS,
    FAIL,
    BLOCKED,
    DELETED,
    CREATED,
    TIMEOUT,
    ERROR,
    NOTFOUND,
    INVALID_PASSWORD,
    INVALID_FIRSTNAME,
    INVALID_LASTNAME,
    INVALID_EMAIL,
    INVALID_PHONE,
    INVALID_ADRESS,
    INVALID_DEVICE,
    INVALID_USERNAME,
    VALIDATED,
    EXITSTS,
    MAXIMUN_NUMBER_OF_USERS_REACHED,
  ];

  static final $core.Map<$core.int, Status> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Status valueOf($core.int value) => _byValue[value];

  const Status._($core.int v, $core.String n) : super(v, n);
}


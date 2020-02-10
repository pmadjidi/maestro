///
//  Generated code. Do not modify.
//  source: register.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/timestamp.pb.dart' as $3;
import 'status.pb.dart' as $0;

class RegisterReq_Address extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegisterReq.Address', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..aOS(1, 'street')
    ..aOS(2, 'city')
    ..aOS(3, 'state')
    ..aOS(4, 'zip')
    ..hasRequiredFields = false
  ;

  RegisterReq_Address._() : super();
  factory RegisterReq_Address() => create();
  factory RegisterReq_Address.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegisterReq_Address.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegisterReq_Address clone() => RegisterReq_Address()..mergeFromMessage(this);
  RegisterReq_Address copyWith(void Function(RegisterReq_Address) updates) => super.copyWith((message) => updates(message as RegisterReq_Address));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegisterReq_Address create() => RegisterReq_Address._();
  RegisterReq_Address createEmptyInstance() => create();
  static $pb.PbList<RegisterReq_Address> createRepeated() => $pb.PbList<RegisterReq_Address>();
  @$core.pragma('dart2js:noInline')
  static RegisterReq_Address getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RegisterReq_Address>(create);
  static RegisterReq_Address _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get street => $_getSZ(0);
  @$pb.TagNumber(1)
  set street($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasStreet() => $_has(0);
  @$pb.TagNumber(1)
  void clearStreet() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get city => $_getSZ(1);
  @$pb.TagNumber(2)
  set city($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCity() => $_has(1);
  @$pb.TagNumber(2)
  void clearCity() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get state => $_getSZ(2);
  @$pb.TagNumber(3)
  set state($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasState() => $_has(2);
  @$pb.TagNumber(3)
  void clearState() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get zip => $_getSZ(3);
  @$pb.TagNumber(4)
  set zip($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasZip() => $_has(3);
  @$pb.TagNumber(4)
  void clearZip() => clearField(4);
}

class RegisterReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegisterReq', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..aOS(1, 'userName', protoName: 'userName')
    ..a<$core.List<$core.int>>(2, 'passWord', $pb.PbFieldType.OY, protoName: 'passWord')
    ..aOS(3, 'firstName', protoName: 'firstName')
    ..aOS(4, 'lastName', protoName: 'lastName')
    ..aOS(5, 'email')
    ..aOS(6, 'phone')
    ..aOM<RegisterReq_Address>(7, 'address', subBuilder: RegisterReq_Address.create)
    ..aOS(8, 'device')
    ..aOM<$3.Timestamp>(9, 'timeName', subBuilder: $3.Timestamp.create)
    ..aOS(10, 'AppName', protoName: 'AppName')
    ..hasRequiredFields = false
  ;

  RegisterReq._() : super();
  factory RegisterReq() => create();
  factory RegisterReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegisterReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegisterReq clone() => RegisterReq()..mergeFromMessage(this);
  RegisterReq copyWith(void Function(RegisterReq) updates) => super.copyWith((message) => updates(message as RegisterReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegisterReq create() => RegisterReq._();
  RegisterReq createEmptyInstance() => create();
  static $pb.PbList<RegisterReq> createRepeated() => $pb.PbList<RegisterReq>();
  @$core.pragma('dart2js:noInline')
  static RegisterReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RegisterReq>(create);
  static RegisterReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userName => $_getSZ(0);
  @$pb.TagNumber(1)
  set userName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserName() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserName() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get passWord => $_getN(1);
  @$pb.TagNumber(2)
  set passWord($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPassWord() => $_has(1);
  @$pb.TagNumber(2)
  void clearPassWord() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get firstName => $_getSZ(2);
  @$pb.TagNumber(3)
  set firstName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasFirstName() => $_has(2);
  @$pb.TagNumber(3)
  void clearFirstName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get lastName => $_getSZ(3);
  @$pb.TagNumber(4)
  set lastName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasLastName() => $_has(3);
  @$pb.TagNumber(4)
  void clearLastName() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get email => $_getSZ(4);
  @$pb.TagNumber(5)
  set email($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasEmail() => $_has(4);
  @$pb.TagNumber(5)
  void clearEmail() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get phone => $_getSZ(5);
  @$pb.TagNumber(6)
  set phone($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasPhone() => $_has(5);
  @$pb.TagNumber(6)
  void clearPhone() => clearField(6);

  @$pb.TagNumber(7)
  RegisterReq_Address get address => $_getN(6);
  @$pb.TagNumber(7)
  set address(RegisterReq_Address v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasAddress() => $_has(6);
  @$pb.TagNumber(7)
  void clearAddress() => clearField(7);
  @$pb.TagNumber(7)
  RegisterReq_Address ensureAddress() => $_ensure(6);

  @$pb.TagNumber(8)
  $core.String get device => $_getSZ(7);
  @$pb.TagNumber(8)
  set device($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasDevice() => $_has(7);
  @$pb.TagNumber(8)
  void clearDevice() => clearField(8);

  @$pb.TagNumber(9)
  $3.Timestamp get timeName => $_getN(8);
  @$pb.TagNumber(9)
  set timeName($3.Timestamp v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasTimeName() => $_has(8);
  @$pb.TagNumber(9)
  void clearTimeName() => clearField(9);
  @$pb.TagNumber(9)
  $3.Timestamp ensureTimeName() => $_ensure(8);

  @$pb.TagNumber(10)
  $core.String get appName => $_getSZ(9);
  @$pb.TagNumber(10)
  set appName($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasAppName() => $_has(9);
  @$pb.TagNumber(10)
  void clearAppName() => clearField(10);
}

class RegisterApi {
  $pb.RpcClient _client;
  RegisterApi(this._client);

  $async.Future<$0.Empty> register($pb.ClientContext ctx, RegisterReq request) {
    var emptyResponse = $0.Empty();
    return _client.invoke<$0.Empty>(ctx, 'Register', 'Register', request, emptyResponse);
  }
}


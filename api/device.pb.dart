///
//  Generated code. Do not modify.
//  source: device.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'status.pbenum.dart' as $0;

class DeviceReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DeviceReq', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..aOS(1, 'deviceId', protoName: 'deviceId')
    ..aOS(2, 'deviceKey', protoName: 'deviceKey')
    ..hasRequiredFields = false
  ;

  DeviceReq._() : super();
  factory DeviceReq() => create();
  factory DeviceReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeviceReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DeviceReq clone() => DeviceReq()..mergeFromMessage(this);
  DeviceReq copyWith(void Function(DeviceReq) updates) => super.copyWith((message) => updates(message as DeviceReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeviceReq create() => DeviceReq._();
  DeviceReq createEmptyInstance() => create();
  static $pb.PbList<DeviceReq> createRepeated() => $pb.PbList<DeviceReq>();
  @$core.pragma('dart2js:noInline')
  static DeviceReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeviceReq>(create);
  static DeviceReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get deviceId => $_getSZ(0);
  @$pb.TagNumber(1)
  set deviceId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDeviceId() => $_has(0);
  @$pb.TagNumber(1)
  void clearDeviceId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get deviceKey => $_getSZ(1);
  @$pb.TagNumber(2)
  set deviceKey($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDeviceKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearDeviceKey() => clearField(2);
}

class DeviceResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DeviceResp', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..e<$0.Status>(1, 'status', $pb.PbFieldType.OE, defaultOrMaker: $0.Status.SUCCESS, valueOf: $0.Status.valueOf, enumValues: $0.Status.values)
    ..hasRequiredFields = false
  ;

  DeviceResp._() : super();
  factory DeviceResp() => create();
  factory DeviceResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeviceResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DeviceResp clone() => DeviceResp()..mergeFromMessage(this);
  DeviceResp copyWith(void Function(DeviceResp) updates) => super.copyWith((message) => updates(message as DeviceResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeviceResp create() => DeviceResp._();
  DeviceResp createEmptyInstance() => create();
  static $pb.PbList<DeviceResp> createRepeated() => $pb.PbList<DeviceResp>();
  @$core.pragma('dart2js:noInline')
  static DeviceResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeviceResp>(create);
  static DeviceResp _defaultInstance;

  @$pb.TagNumber(1)
  $0.Status get status => $_getN(0);
  @$pb.TagNumber(1)
  set status($0.Status v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(1)
  void clearStatus() => clearField(1);
}

class DeviceApi {
  $pb.RpcClient _client;
  DeviceApi(this._client);

  $async.Future<DeviceResp> register($pb.ClientContext ctx, DeviceReq request) {
    var emptyResponse = DeviceResp();
    return _client.invoke<DeviceResp>(ctx, 'Device', 'Register', request, emptyResponse);
  }
}


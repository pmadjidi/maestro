///
//  Generated code. Do not modify.
//  source: message.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/timestamp.pb.dart' as $2;

import 'status.pbenum.dart' as $0;

class MsgReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('MsgReq', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..aOS(1, 'text')
    ..a<$core.List<$core.int>>(2, 'pic', $pb.PbFieldType.OY)
    ..aOS(3, 'parentId', protoName: 'parentId')
    ..aOS(4, 'topic')
    ..aOM<$2.Timestamp>(5, 'timeName', subBuilder: $2.Timestamp.create)
    ..e<$0.Status>(6, 'status', $pb.PbFieldType.OE, defaultOrMaker: $0.Status.SUCCESS, valueOf: $0.Status.valueOf, enumValues: $0.Status.values)
    ..aOS(7, 'uuid')
    ..hasRequiredFields = false
  ;

  MsgReq._() : super();
  factory MsgReq() => create();
  factory MsgReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  MsgReq clone() => MsgReq()..mergeFromMessage(this);
  MsgReq copyWith(void Function(MsgReq) updates) => super.copyWith((message) => updates(message as MsgReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgReq create() => MsgReq._();
  MsgReq createEmptyInstance() => create();
  static $pb.PbList<MsgReq> createRepeated() => $pb.PbList<MsgReq>();
  @$core.pragma('dart2js:noInline')
  static MsgReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgReq>(create);
  static MsgReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get text => $_getSZ(0);
  @$pb.TagNumber(1)
  set text($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasText() => $_has(0);
  @$pb.TagNumber(1)
  void clearText() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get pic => $_getN(1);
  @$pb.TagNumber(2)
  set pic($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPic() => $_has(1);
  @$pb.TagNumber(2)
  void clearPic() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get parentId => $_getSZ(2);
  @$pb.TagNumber(3)
  set parentId($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasParentId() => $_has(2);
  @$pb.TagNumber(3)
  void clearParentId() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get topic => $_getSZ(3);
  @$pb.TagNumber(4)
  set topic($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTopic() => $_has(3);
  @$pb.TagNumber(4)
  void clearTopic() => clearField(4);

  @$pb.TagNumber(5)
  $2.Timestamp get timeName => $_getN(4);
  @$pb.TagNumber(5)
  set timeName($2.Timestamp v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasTimeName() => $_has(4);
  @$pb.TagNumber(5)
  void clearTimeName() => clearField(5);
  @$pb.TagNumber(5)
  $2.Timestamp ensureTimeName() => $_ensure(4);

  @$pb.TagNumber(6)
  $0.Status get status => $_getN(5);
  @$pb.TagNumber(6)
  set status($0.Status v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasStatus() => $_has(5);
  @$pb.TagNumber(6)
  void clearStatus() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get uuid => $_getSZ(6);
  @$pb.TagNumber(7)
  set uuid($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasUuid() => $_has(6);
  @$pb.TagNumber(7)
  void clearUuid() => clearField(7);
}

class MsgResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('MsgResp', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..e<$0.Status>(1, 'status', $pb.PbFieldType.OE, defaultOrMaker: $0.Status.SUCCESS, valueOf: $0.Status.valueOf, enumValues: $0.Status.values)
    ..aOS(2, 'uuid')
    ..hasRequiredFields = false
  ;

  MsgResp._() : super();
  factory MsgResp() => create();
  factory MsgResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  MsgResp clone() => MsgResp()..mergeFromMessage(this);
  MsgResp copyWith(void Function(MsgResp) updates) => super.copyWith((message) => updates(message as MsgResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgResp create() => MsgResp._();
  MsgResp createEmptyInstance() => create();
  static $pb.PbList<MsgResp> createRepeated() => $pb.PbList<MsgResp>();
  @$core.pragma('dart2js:noInline')
  static MsgResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgResp>(create);
  static MsgResp _defaultInstance;

  @$pb.TagNumber(1)
  $0.Status get status => $_getN(0);
  @$pb.TagNumber(1)
  set status($0.Status v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(1)
  void clearStatus() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get uuid => $_getSZ(1);
  @$pb.TagNumber(2)
  set uuid($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasUuid() => $_has(1);
  @$pb.TagNumber(2)
  void clearUuid() => clearField(2);
}

class MsgApi {
  $pb.RpcClient _client;
  MsgApi(this._client);

  $async.Future<MsgResp> put($pb.ClientContext ctx, MsgReq request) {
    var emptyResponse = MsgResp();
    return _client.invoke<MsgResp>(ctx, 'Msg', 'put', request, emptyResponse);
  }
}


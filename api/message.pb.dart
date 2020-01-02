///
//  Generated code. Do not modify.
//  source: message.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/timestamp.pb.dart' as $3;

import 'status.pbenum.dart' as $0;

class MsgReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('MsgReq', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..aOS(1, 'id')
    ..pPS(2, 'text')
    ..p<$core.List<$core.int>>(3, 'pic', $pb.PbFieldType.PY)
    ..aOS(4, 'parentId', protoName: 'parentId')
    ..aOS(5, 'topic')
    ..aOM<$3.Timestamp>(6, 'timeName', subBuilder: $3.Timestamp.create)
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
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.String> get text => $_getList(1);

  @$pb.TagNumber(3)
  $core.List<$core.List<$core.int>> get pic => $_getList(2);

  @$pb.TagNumber(4)
  $core.String get parentId => $_getSZ(3);
  @$pb.TagNumber(4)
  set parentId($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasParentId() => $_has(3);
  @$pb.TagNumber(4)
  void clearParentId() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get topic => $_getSZ(4);
  @$pb.TagNumber(5)
  set topic($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasTopic() => $_has(4);
  @$pb.TagNumber(5)
  void clearTopic() => clearField(5);

  @$pb.TagNumber(6)
  $3.Timestamp get timeName => $_getN(5);
  @$pb.TagNumber(6)
  set timeName($3.Timestamp v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasTimeName() => $_has(5);
  @$pb.TagNumber(6)
  void clearTimeName() => clearField(6);
  @$pb.TagNumber(6)
  $3.Timestamp ensureTimeName() => $_ensure(5);
}

class MsgResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('MsgResp', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..aOS(1, 'id')
    ..e<$0.Status>(2, 'status', $pb.PbFieldType.OE, defaultOrMaker: $0.Status.SUCCESS, valueOf: $0.Status.valueOf, enumValues: $0.Status.values)
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
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $0.Status get status => $_getN(1);
  @$pb.TagNumber(2)
  set status($0.Status v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasStatus() => $_has(1);
  @$pb.TagNumber(2)
  void clearStatus() => clearField(2);
}

class MaestroApi {
  $pb.RpcClient _client;
  MaestroApi(this._client);

  $async.Future<MsgResp> chatt($pb.ClientContext ctx, MsgReq request) {
    var emptyResponse = MsgResp();
    return _client.invoke<MsgResp>(ctx, 'Maestro', 'Chatt', request, emptyResponse);
  }
}


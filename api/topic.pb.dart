///
//  Generated code. Do not modify.
//  source: topic.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'status.pbenum.dart' as $0;

class TopicReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('TopicReq', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..pPS(1, 'topic')
    ..hasRequiredFields = false
  ;

  TopicReq._() : super();
  factory TopicReq() => create();
  factory TopicReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TopicReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  TopicReq clone() => TopicReq()..mergeFromMessage(this);
  TopicReq copyWith(void Function(TopicReq) updates) => super.copyWith((message) => updates(message as TopicReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TopicReq create() => TopicReq._();
  TopicReq createEmptyInstance() => create();
  static $pb.PbList<TopicReq> createRepeated() => $pb.PbList<TopicReq>();
  @$core.pragma('dart2js:noInline')
  static TopicReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TopicReq>(create);
  static TopicReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get topic => $_getList(0);
}

class TopicResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('TopicResp', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..aOS(1, 'id')
    ..e<$0.Status>(2, 'status', $pb.PbFieldType.OE, defaultOrMaker: $0.Status.SUCCESS, valueOf: $0.Status.valueOf, enumValues: $0.Status.values)
    ..hasRequiredFields = false
  ;

  TopicResp._() : super();
  factory TopicResp() => create();
  factory TopicResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TopicResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  TopicResp clone() => TopicResp()..mergeFromMessage(this);
  TopicResp copyWith(void Function(TopicResp) updates) => super.copyWith((message) => updates(message as TopicResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TopicResp create() => TopicResp._();
  TopicResp createEmptyInstance() => create();
  static $pb.PbList<TopicResp> createRepeated() => $pb.PbList<TopicResp>();
  @$core.pragma('dart2js:noInline')
  static TopicResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TopicResp>(create);
  static TopicResp _defaultInstance;

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

class TopicApi {
  $pb.RpcClient _client;
  TopicApi(this._client);

  $async.Future<TopicResp> subscribe($pb.ClientContext ctx, TopicReq request) {
    var emptyResponse = TopicResp();
    return _client.invoke<TopicResp>(ctx, 'Topic', 'subscribe', request, emptyResponse);
  }
  $async.Future<TopicResp> unsubscribe($pb.ClientContext ctx, TopicReq request) {
    var emptyResponse = TopicResp();
    return _client.invoke<TopicResp>(ctx, 'Topic', 'unsubscribe', request, emptyResponse);
  }
}


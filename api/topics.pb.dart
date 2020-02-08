///
//  Generated code. Do not modify.
//  source: topics.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'register.pb.dart' as $3;

import 'status.pbenum.dart' as $0;

class Topic extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Topic', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..aOS(1, 'id')
    ..aOS(2, 'tag')
    ..e<$0.Status>(3, 'status', $pb.PbFieldType.OE, defaultOrMaker: $0.Status.SUCCESS, valueOf: $0.Status.valueOf, enumValues: $0.Status.values)
    ..hasRequiredFields = false
  ;

  Topic._() : super();
  factory Topic() => create();
  factory Topic.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Topic.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Topic clone() => Topic()..mergeFromMessage(this);
  Topic copyWith(void Function(Topic) updates) => super.copyWith((message) => updates(message as Topic));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Topic create() => Topic._();
  Topic createEmptyInstance() => create();
  static $pb.PbList<Topic> createRepeated() => $pb.PbList<Topic>();
  @$core.pragma('dart2js:noInline')
  static Topic getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Topic>(create);
  static Topic _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get tag => $_getSZ(1);
  @$pb.TagNumber(2)
  set tag($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTag() => $_has(1);
  @$pb.TagNumber(2)
  void clearTag() => clearField(2);

  @$pb.TagNumber(3)
  $0.Status get status => $_getN(2);
  @$pb.TagNumber(3)
  set status($0.Status v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasStatus() => $_has(2);
  @$pb.TagNumber(3)
  void clearStatus() => clearField(3);
}

class TopicReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('TopicReq', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..pc<Topic>(1, 'list', $pb.PbFieldType.PM, subBuilder: Topic.create)
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
  $core.List<Topic> get list => $_getList(0);
}

class TopicResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('TopicResp', package: const $pb.PackageName('api'), createEmptyInstance: create)
    ..pc<Topic>(1, 'list', $pb.PbFieldType.PM, subBuilder: Topic.create)
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
  $core.List<Topic> get list => $_getList(0);
}

class SubscriptionsApi {
  $pb.RpcClient _client;
  SubscriptionsApi(this._client);

  $async.Future<TopicResp> sub($pb.ClientContext ctx, TopicReq request) {
    var emptyResponse = TopicResp();
    return _client.invoke<TopicResp>(ctx, 'Subscriptions', 'sub', request, emptyResponse);
  }
  $async.Future<TopicResp> unsub($pb.ClientContext ctx, TopicReq request) {
    var emptyResponse = TopicResp();
    return _client.invoke<TopicResp>(ctx, 'Subscriptions', 'unsub', request, emptyResponse);
  }
  $async.Future<TopicResp> list($pb.ClientContext ctx, $3.Empty request) {
    var emptyResponse = TopicResp();
    return _client.invoke<TopicResp>(ctx, 'Subscriptions', 'list', request, emptyResponse);
  }
}


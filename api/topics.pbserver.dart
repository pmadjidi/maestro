///
//  Generated code. Do not modify.
//  source: topics.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'topics.pb.dart' as $6;
import 'status.pb.dart' as $0;
import 'topics.pbjson.dart';

export 'topics.pb.dart';

abstract class SubscriptionsServiceBase extends $pb.GeneratedService {
  $async.Future<$6.TopicResp> sub($pb.ServerContext ctx, $6.TopicReq request);
  $async.Future<$6.TopicResp> unsub($pb.ServerContext ctx, $6.TopicReq request);
  $async.Future<$6.TopicResp> list($pb.ServerContext ctx, $0.Empty request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'sub': return $6.TopicReq();
      case 'unsub': return $6.TopicReq();
      case 'list': return $0.Empty();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'sub': return this.sub(ctx, request);
      case 'unsub': return this.unsub(ctx, request);
      case 'list': return this.list(ctx, request);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => SubscriptionsServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => SubscriptionsServiceBase$messageJson;
}


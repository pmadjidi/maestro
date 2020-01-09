///
//  Generated code. Do not modify.
//  source: topic.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'topic.pb.dart' as $6;
import 'topic.pbjson.dart';

export 'topic.pb.dart';

abstract class TopicServiceBase extends $pb.GeneratedService {
  $async.Future<$6.TopicResp> subscribe($pb.ServerContext ctx, $6.TopicReq request);
  $async.Future<$6.TopicResp> unsubscribe($pb.ServerContext ctx, $6.TopicReq request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'subscribe': return $6.TopicReq();
      case 'unsubscribe': return $6.TopicReq();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'subscribe': return this.subscribe(ctx, request);
      case 'unsubscribe': return this.unsubscribe(ctx, request);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => TopicServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => TopicServiceBase$messageJson;
}


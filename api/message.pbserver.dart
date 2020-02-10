///
//  Generated code. Do not modify.
//  source: message.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'message.pb.dart' as $4;
import 'status.pb.dart' as $0;
import 'message.pbjson.dart';

export 'message.pb.dart';

abstract class MsgServiceBase extends $pb.GeneratedService {
  $async.Future<$4.MsgResp> put($pb.ServerContext ctx, $4.MsgReq request);
  $async.Future<$4.MsgReq> timeLine($pb.ServerContext ctx, $0.Empty request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'put': return $4.MsgReq();
      case 'timeLine': return $0.Empty();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'put': return this.put(ctx, request);
      case 'timeLine': return this.timeLine(ctx, request);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => MsgServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => MsgServiceBase$messageJson;
}


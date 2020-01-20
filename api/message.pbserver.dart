///
//  Generated code. Do not modify.
//  source: message.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'message.pb.dart' as $5;
import 'message.pbjson.dart';

export 'message.pb.dart';

abstract class MessageServiceBase extends $pb.GeneratedService {
  $async.Future<$5.MsgReq> msg($pb.ServerContext ctx, $5.MsgReq request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Msg': return $5.MsgReq();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Msg': return this.msg(ctx, request);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => MessageServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => MessageServiceBase$messageJson;
}


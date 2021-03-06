///
//  Generated code. Do not modify.
//  source: login.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'login.pb.dart' as $2;
import 'status.pb.dart' as $0;
import 'login.pbjson.dart';

export 'login.pb.dart';

abstract class LoginServiceBase extends $pb.GeneratedService {
  $async.Future<$0.Empty> authenticate($pb.ServerContext ctx, $2.LoginReq request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Authenticate': return $2.LoginReq();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Authenticate': return this.authenticate(ctx, request);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => LoginServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => LoginServiceBase$messageJson;
}


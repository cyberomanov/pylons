// Mocks generated by Mockito 5.2.0 from annotations
// in pylons_wallet/test/unit_testing/services/data_stores/remote_data_store_test.dart.
// Do not manually edit this file.

import 'dart:async' as _i4;

import 'package:firebase_app_check/firebase_app_check.dart' as _i3;
import 'package:firebase_core/firebase_core.dart' as _i2;
import 'package:mockito/mockito.dart' as _i1;

// ignore_for_file: type=lint
// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types
// ignore_for_file: invalid_override,annotate_overrides,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

class _FakeFirebaseApp_0 extends _i1.Fake implements _i2.FirebaseApp {}

/// A class which mocks [FirebaseAppCheck].
///
/// See the documentation for Mockito's code generation for more information.
class MockFirebaseAppCheck extends _i1.Mock implements _i3.FirebaseAppCheck {
  MockFirebaseAppCheck() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i2.FirebaseApp get app => (super.noSuchMethod(Invocation.getter(#app),
      returnValue: _FakeFirebaseApp_0()) as _i2.FirebaseApp);
  @override
  set app(_i2.FirebaseApp? _app) =>
      super.noSuchMethod(Invocation.setter(#app, _app),
          returnValueForMissingStub: null);
  @override
  _i4.Stream<String?> get onTokenChange =>
      (super.noSuchMethod(Invocation.getter(#onTokenChange),
          returnValue: Stream<String?>.empty()) as _i4.Stream<String?>);
  @override
  Map<dynamic, dynamic> get pluginConstants =>
      (super.noSuchMethod(Invocation.getter(#pluginConstants),
          returnValue: <dynamic, dynamic>{}) as Map<dynamic, dynamic>);
  @override
  _i4.Future<void> activate({String? webRecaptchaSiteKey}) =>
      (super.noSuchMethod(
          Invocation.method(
              #activate, [], {#webRecaptchaSiteKey: webRecaptchaSiteKey}),
          returnValue: Future<void>.value(),
          returnValueForMissingStub: Future<void>.value()) as _i4.Future<void>);
  @override
  _i4.Future<String?> getToken([bool? forceRefresh]) =>
      (super.noSuchMethod(Invocation.method(#getToken, [forceRefresh]),
          returnValue: Future<String?>.value()) as _i4.Future<String?>);
  @override
  _i4.Future<void> setTokenAutoRefreshEnabled(
          bool? isTokenAutoRefreshEnabled) =>
      (super.noSuchMethod(
          Invocation.method(
              #setTokenAutoRefreshEnabled, [isTokenAutoRefreshEnabled]),
          returnValue: Future<void>.value(),
          returnValueForMissingStub: Future<void>.value()) as _i4.Future<void>);
}

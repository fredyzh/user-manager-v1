// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: service.proto

package com.client.test.pb.grpc;

public interface GetUserReplyOrBuilder extends
    // @@protoc_insertion_point(interface_extends:user.v1.service.GetUserReply)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int64 id = 1;</code>
   * @return The id.
   */
  long getId();

  /**
   * <code>.user.v1.pb.User user = 2;</code>
   * @return Whether the user field is set.
   */
  boolean hasUser();
  /**
   * <code>.user.v1.pb.User user = 2;</code>
   * @return The user.
   */
  User getUser();
  /**
   * <code>.user.v1.pb.User user = 2;</code>
   */
  UserOrBuilder getUserOrBuilder();
}
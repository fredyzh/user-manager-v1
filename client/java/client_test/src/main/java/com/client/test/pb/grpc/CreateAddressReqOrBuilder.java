// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: service.proto

package com.client.test.pb.grpc;

public interface CreateAddressReqOrBuilder extends
    // @@protoc_insertion_point(interface_extends:user.v1.service.CreateAddressReq)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string id = 1;</code>
   * @return The id.
   */
  java.lang.String getId();
  /**
   * <code>string id = 1;</code>
   * @return The bytes for id.
   */
  com.google.protobuf.ByteString
      getIdBytes();

  /**
   * <code>.user.v1.pb.Address address = 2;</code>
   * @return Whether the address field is set.
   */
  boolean hasAddress();
  /**
   * <code>.user.v1.pb.Address address = 2;</code>
   * @return The address.
   */
  Address getAddress();
  /**
   * <code>.user.v1.pb.Address address = 2;</code>
   */
  AddressOrBuilder getAddressOrBuilder();
}

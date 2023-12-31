// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: user.proto

package com.client.test.pb.grpc;

/**
 * Protobuf type {@code user.v1.pb.Address}
 */
public final class Address extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:user.v1.pb.Address)
    AddressOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Address.newBuilder() to construct.
  private Address(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Address() {
    street_ = "";
    city_ = "";
    state_ = "";
    zip_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Address();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return GRPCUserClientTest.internal_static_user_v1_pb_Address_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return GRPCUserClientTest.internal_static_user_v1_pb_Address_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            Address.class, Address.Builder.class);
  }

  public static final int STREET_FIELD_NUMBER = 1;
  @SuppressWarnings("serial")
  private volatile java.lang.Object street_ = "";
  /**
   * <pre>
   * &#64;gotags: bson:"street" json:"street,omitempty" 
   * </pre>
   *
   * <code>string street = 1;</code>
   * @return The street.
   */
  @java.lang.Override
  public java.lang.String getStreet() {
    java.lang.Object ref = street_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      street_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * &#64;gotags: bson:"street" json:"street,omitempty" 
   * </pre>
   *
   * <code>string street = 1;</code>
   * @return The bytes for street.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getStreetBytes() {
    java.lang.Object ref = street_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      street_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int CITY_FIELD_NUMBER = 2;
  @SuppressWarnings("serial")
  private volatile java.lang.Object city_ = "";
  /**
   * <pre>
   * &#64;gotags: bson:"city" json:"city,omitempty" 
   * </pre>
   *
   * <code>string city = 2;</code>
   * @return The city.
   */
  @java.lang.Override
  public java.lang.String getCity() {
    java.lang.Object ref = city_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      city_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * &#64;gotags: bson:"city" json:"city,omitempty" 
   * </pre>
   *
   * <code>string city = 2;</code>
   * @return The bytes for city.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getCityBytes() {
    java.lang.Object ref = city_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      city_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int STATE_FIELD_NUMBER = 3;
  @SuppressWarnings("serial")
  private volatile java.lang.Object state_ = "";
  /**
   * <pre>
   * &#64;gotags: bson:"state" json:"state,omitempty" 
   * </pre>
   *
   * <code>string state = 3;</code>
   * @return The state.
   */
  @java.lang.Override
  public java.lang.String getState() {
    java.lang.Object ref = state_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      state_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * &#64;gotags: bson:"state" json:"state,omitempty" 
   * </pre>
   *
   * <code>string state = 3;</code>
   * @return The bytes for state.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getStateBytes() {
    java.lang.Object ref = state_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      state_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int ZIP_FIELD_NUMBER = 4;
  @SuppressWarnings("serial")
  private volatile java.lang.Object zip_ = "";
  /**
   * <pre>
   * &#64;gotags: bson:"zip" json:"zip,omitempty" 
   * </pre>
   *
   * <code>string zip = 4;</code>
   * @return The zip.
   */
  @java.lang.Override
  public java.lang.String getZip() {
    java.lang.Object ref = zip_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      zip_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * &#64;gotags: bson:"zip" json:"zip,omitempty" 
   * </pre>
   *
   * <code>string zip = 4;</code>
   * @return The bytes for zip.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getZipBytes() {
    java.lang.Object ref = zip_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      zip_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(street_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, street_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(city_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, city_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(state_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 3, state_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(zip_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 4, zip_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(street_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, street_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(city_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, city_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(state_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, state_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(zip_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(4, zip_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof Address)) {
      return super.equals(obj);
    }
    Address other = (Address) obj;

    if (!getStreet()
        .equals(other.getStreet())) return false;
    if (!getCity()
        .equals(other.getCity())) return false;
    if (!getState()
        .equals(other.getState())) return false;
    if (!getZip()
        .equals(other.getZip())) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + STREET_FIELD_NUMBER;
    hash = (53 * hash) + getStreet().hashCode();
    hash = (37 * hash) + CITY_FIELD_NUMBER;
    hash = (53 * hash) + getCity().hashCode();
    hash = (37 * hash) + STATE_FIELD_NUMBER;
    hash = (53 * hash) + getState().hashCode();
    hash = (37 * hash) + ZIP_FIELD_NUMBER;
    hash = (53 * hash) + getZip().hashCode();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static Address parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static Address parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static Address parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static Address parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static Address parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static Address parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static Address parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static Address parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static Address parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static Address parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static Address parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static Address parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(Address prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code user.v1.pb.Address}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:user.v1.pb.Address)
      AddressOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return GRPCUserClientTest.internal_static_user_v1_pb_Address_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return GRPCUserClientTest.internal_static_user_v1_pb_Address_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              Address.class, Address.Builder.class);
    }

    // Construct using grpc.Address.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      street_ = "";
      city_ = "";
      state_ = "";
      zip_ = "";
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return GRPCUserClientTest.internal_static_user_v1_pb_Address_descriptor;
    }

    @java.lang.Override
    public Address getDefaultInstanceForType() {
      return Address.getDefaultInstance();
    }

    @java.lang.Override
    public Address build() {
      Address result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public Address buildPartial() {
      Address result = new Address(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(Address result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.street_ = street_;
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.city_ = city_;
      }
      if (((from_bitField0_ & 0x00000004) != 0)) {
        result.state_ = state_;
      }
      if (((from_bitField0_ & 0x00000008) != 0)) {
        result.zip_ = zip_;
      }
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof Address) {
        return mergeFrom((Address)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(Address other) {
      if (other == Address.getDefaultInstance()) return this;
      if (!other.getStreet().isEmpty()) {
        street_ = other.street_;
        bitField0_ |= 0x00000001;
        onChanged();
      }
      if (!other.getCity().isEmpty()) {
        city_ = other.city_;
        bitField0_ |= 0x00000002;
        onChanged();
      }
      if (!other.getState().isEmpty()) {
        state_ = other.state_;
        bitField0_ |= 0x00000004;
        onChanged();
      }
      if (!other.getZip().isEmpty()) {
        zip_ = other.zip_;
        bitField0_ |= 0x00000008;
        onChanged();
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              street_ = input.readStringRequireUtf8();
              bitField0_ |= 0x00000001;
              break;
            } // case 10
            case 18: {
              city_ = input.readStringRequireUtf8();
              bitField0_ |= 0x00000002;
              break;
            } // case 18
            case 26: {
              state_ = input.readStringRequireUtf8();
              bitField0_ |= 0x00000004;
              break;
            } // case 26
            case 34: {
              zip_ = input.readStringRequireUtf8();
              bitField0_ |= 0x00000008;
              break;
            } // case 34
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }
    private int bitField0_;

    private java.lang.Object street_ = "";
    /**
     * <pre>
     * &#64;gotags: bson:"street" json:"street,omitempty" 
     * </pre>
     *
     * <code>string street = 1;</code>
     * @return The street.
     */
    public java.lang.String getStreet() {
      java.lang.Object ref = street_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        street_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     * &#64;gotags: bson:"street" json:"street,omitempty" 
     * </pre>
     *
     * <code>string street = 1;</code>
     * @return The bytes for street.
     */
    public com.google.protobuf.ByteString
        getStreetBytes() {
      java.lang.Object ref = street_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        street_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * &#64;gotags: bson:"street" json:"street,omitempty" 
     * </pre>
     *
     * <code>string street = 1;</code>
     * @param value The street to set.
     * @return This builder for chaining.
     */
    public Builder setStreet(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      street_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * &#64;gotags: bson:"street" json:"street,omitempty" 
     * </pre>
     *
     * <code>string street = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearStreet() {
      street_ = getDefaultInstance().getStreet();
      bitField0_ = (bitField0_ & ~0x00000001);
      onChanged();
      return this;
    }
    /**
     * <pre>
     * &#64;gotags: bson:"street" json:"street,omitempty" 
     * </pre>
     *
     * <code>string street = 1;</code>
     * @param value The bytes for street to set.
     * @return This builder for chaining.
     */
    public Builder setStreetBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      street_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }

    private java.lang.Object city_ = "";
    /**
     * <pre>
     * &#64;gotags: bson:"city" json:"city,omitempty" 
     * </pre>
     *
     * <code>string city = 2;</code>
     * @return The city.
     */
    public java.lang.String getCity() {
      java.lang.Object ref = city_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        city_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     * &#64;gotags: bson:"city" json:"city,omitempty" 
     * </pre>
     *
     * <code>string city = 2;</code>
     * @return The bytes for city.
     */
    public com.google.protobuf.ByteString
        getCityBytes() {
      java.lang.Object ref = city_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        city_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * &#64;gotags: bson:"city" json:"city,omitempty" 
     * </pre>
     *
     * <code>string city = 2;</code>
     * @param value The city to set.
     * @return This builder for chaining.
     */
    public Builder setCity(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      city_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * &#64;gotags: bson:"city" json:"city,omitempty" 
     * </pre>
     *
     * <code>string city = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearCity() {
      city_ = getDefaultInstance().getCity();
      bitField0_ = (bitField0_ & ~0x00000002);
      onChanged();
      return this;
    }
    /**
     * <pre>
     * &#64;gotags: bson:"city" json:"city,omitempty" 
     * </pre>
     *
     * <code>string city = 2;</code>
     * @param value The bytes for city to set.
     * @return This builder for chaining.
     */
    public Builder setCityBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      city_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }

    private java.lang.Object state_ = "";
    /**
     * <pre>
     * &#64;gotags: bson:"state" json:"state,omitempty" 
     * </pre>
     *
     * <code>string state = 3;</code>
     * @return The state.
     */
    public java.lang.String getState() {
      java.lang.Object ref = state_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        state_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     * &#64;gotags: bson:"state" json:"state,omitempty" 
     * </pre>
     *
     * <code>string state = 3;</code>
     * @return The bytes for state.
     */
    public com.google.protobuf.ByteString
        getStateBytes() {
      java.lang.Object ref = state_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        state_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * &#64;gotags: bson:"state" json:"state,omitempty" 
     * </pre>
     *
     * <code>string state = 3;</code>
     * @param value The state to set.
     * @return This builder for chaining.
     */
    public Builder setState(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      state_ = value;
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * &#64;gotags: bson:"state" json:"state,omitempty" 
     * </pre>
     *
     * <code>string state = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearState() {
      state_ = getDefaultInstance().getState();
      bitField0_ = (bitField0_ & ~0x00000004);
      onChanged();
      return this;
    }
    /**
     * <pre>
     * &#64;gotags: bson:"state" json:"state,omitempty" 
     * </pre>
     *
     * <code>string state = 3;</code>
     * @param value The bytes for state to set.
     * @return This builder for chaining.
     */
    public Builder setStateBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      state_ = value;
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }

    private java.lang.Object zip_ = "";
    /**
     * <pre>
     * &#64;gotags: bson:"zip" json:"zip,omitempty" 
     * </pre>
     *
     * <code>string zip = 4;</code>
     * @return The zip.
     */
    public java.lang.String getZip() {
      java.lang.Object ref = zip_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        zip_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     * &#64;gotags: bson:"zip" json:"zip,omitempty" 
     * </pre>
     *
     * <code>string zip = 4;</code>
     * @return The bytes for zip.
     */
    public com.google.protobuf.ByteString
        getZipBytes() {
      java.lang.Object ref = zip_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        zip_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * &#64;gotags: bson:"zip" json:"zip,omitempty" 
     * </pre>
     *
     * <code>string zip = 4;</code>
     * @param value The zip to set.
     * @return This builder for chaining.
     */
    public Builder setZip(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      zip_ = value;
      bitField0_ |= 0x00000008;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * &#64;gotags: bson:"zip" json:"zip,omitempty" 
     * </pre>
     *
     * <code>string zip = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearZip() {
      zip_ = getDefaultInstance().getZip();
      bitField0_ = (bitField0_ & ~0x00000008);
      onChanged();
      return this;
    }
    /**
     * <pre>
     * &#64;gotags: bson:"zip" json:"zip,omitempty" 
     * </pre>
     *
     * <code>string zip = 4;</code>
     * @param value The bytes for zip to set.
     * @return This builder for chaining.
     */
    public Builder setZipBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      zip_ = value;
      bitField0_ |= 0x00000008;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:user.v1.pb.Address)
  }

  // @@protoc_insertion_point(class_scope:user.v1.pb.Address)
  private static final Address DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new Address();
  }

  public static Address getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Address>
      PARSER = new com.google.protobuf.AbstractParser<Address>() {
    @java.lang.Override
    public Address parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<Address> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Address> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public Address getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}


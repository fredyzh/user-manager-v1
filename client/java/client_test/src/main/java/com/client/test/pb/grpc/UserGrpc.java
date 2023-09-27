package com.client.test.pb.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.58.0)",
    comments = "Source: service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class UserGrpc {

  private UserGrpc() {}

  public static final java.lang.String SERVICE_NAME = "user.v1.service.User";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      HealthReply> getHealthMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Health",
      requestType = com.google.protobuf.Empty.class,
      responseType = HealthReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      HealthReply> getHealthMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, HealthReply> getHealthMethod;
    if ((getHealthMethod = UserGrpc.getHealthMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getHealthMethod = UserGrpc.getHealthMethod) == null) {
        	UserGrpc.getHealthMethod = getHealthMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, HealthReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Health"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  HealthReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("Health"))
              .build();
        }
      }
    }
    return getHealthMethod;
  }

  private static volatile io.grpc.MethodDescriptor<GetUserReq,
      GetUserReply> getGetUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUser",
      requestType = GetUserReq.class,
      responseType = GetUserReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<GetUserReq,
      GetUserReply> getGetUserMethod() {
    io.grpc.MethodDescriptor<GetUserReq, GetUserReply> getGetUserMethod;
    if ((getGetUserMethod = UserGrpc.getGetUserMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getGetUserMethod = UserGrpc.getGetUserMethod) == null) {
        	UserGrpc.getGetUserMethod = getGetUserMethod =
              io.grpc.MethodDescriptor.<GetUserReq, GetUserReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  GetUserReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  GetUserReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("GetUser"))
              .build();
        }
      }
    }
    return getGetUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<GetUserByUsernameReq,
      GetUserByUsernameReply> getGetUserByUsernameMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserByUsername",
      requestType = GetUserByUsernameReq.class,
      responseType = GetUserByUsernameReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<GetUserByUsernameReq,
      GetUserByUsernameReply> getGetUserByUsernameMethod() {
    io.grpc.MethodDescriptor<GetUserByUsernameReq, GetUserByUsernameReply> getGetUserByUsernameMethod;
    if ((getGetUserByUsernameMethod = UserGrpc.getGetUserByUsernameMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getGetUserByUsernameMethod = UserGrpc.getGetUserByUsernameMethod) == null) {
        	UserGrpc.getGetUserByUsernameMethod = getGetUserByUsernameMethod =
              io.grpc.MethodDescriptor.<GetUserByUsernameReq, GetUserByUsernameReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserByUsername"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  GetUserByUsernameReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  GetUserByUsernameReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("GetUserByUsername"))
              .build();
        }
      }
    }
    return getGetUserByUsernameMethod;
  }

  private static volatile io.grpc.MethodDescriptor<ResetUserPasswordReq,
      ResetUserPasswordReply> getResetUserPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ResetUserPassword",
      requestType = ResetUserPasswordReq.class,
      responseType = ResetUserPasswordReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<ResetUserPasswordReq,
      ResetUserPasswordReply> getResetUserPasswordMethod() {
    io.grpc.MethodDescriptor<ResetUserPasswordReq, ResetUserPasswordReply> getResetUserPasswordMethod;
    if ((getResetUserPasswordMethod = UserGrpc.getResetUserPasswordMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getResetUserPasswordMethod = UserGrpc.getResetUserPasswordMethod) == null) {
          UserGrpc.getResetUserPasswordMethod = getResetUserPasswordMethod =
              io.grpc.MethodDescriptor.<ResetUserPasswordReq, ResetUserPasswordReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ResetUserPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  ResetUserPasswordReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  ResetUserPasswordReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("ResetUserPassword"))
              .build();
        }
      }
    }
    return getResetUserPasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<CreateUserReq,
      CreateUserReply> getCreateUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateUser",
      requestType = CreateUserReq.class,
      responseType = CreateUserReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<CreateUserReq,
      CreateUserReply> getCreateUserMethod() {
    io.grpc.MethodDescriptor<CreateUserReq, CreateUserReply> getCreateUserMethod;
    if ((getCreateUserMethod = UserGrpc.getCreateUserMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getCreateUserMethod = UserGrpc.getCreateUserMethod) == null) {
          UserGrpc.getCreateUserMethod = getCreateUserMethod =
              io.grpc.MethodDescriptor.<CreateUserReq, CreateUserReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  CreateUserReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  CreateUserReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("CreateUser"))
              .build();
        }
      }
    }
    return getCreateUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<ListUserReq,
      ListUserReply> getListUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListUser",
      requestType = ListUserReq.class,
      responseType = ListUserReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<ListUserReq,
      ListUserReply> getListUserMethod() {
    io.grpc.MethodDescriptor<ListUserReq, ListUserReply> getListUserMethod;
    if ((getListUserMethod = UserGrpc.getListUserMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getListUserMethod = UserGrpc.getListUserMethod) == null) {
          UserGrpc.getListUserMethod = getListUserMethod =
              io.grpc.MethodDescriptor.<ListUserReq, ListUserReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  ListUserReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  ListUserReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("ListUser"))
              .build();
        }
      }
    }
    return getListUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<VerifyPasswordReq,
      VerifyPasswordReply> getVerifyPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "VerifyPassword",
      requestType = VerifyPasswordReq.class,
      responseType = VerifyPasswordReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<VerifyPasswordReq,
      VerifyPasswordReply> getVerifyPasswordMethod() {
    io.grpc.MethodDescriptor<VerifyPasswordReq, VerifyPasswordReply> getVerifyPasswordMethod;
    if ((getVerifyPasswordMethod = UserGrpc.getVerifyPasswordMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getVerifyPasswordMethod = UserGrpc.getVerifyPasswordMethod) == null) {
          UserGrpc.getVerifyPasswordMethod = getVerifyPasswordMethod =
              io.grpc.MethodDescriptor.<VerifyPasswordReq, VerifyPasswordReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "VerifyPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  VerifyPasswordReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  VerifyPasswordReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("VerifyPassword"))
              .build();
        }
      }
    }
    return getVerifyPasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<CreateAddressReq,
      CreateAddressReply> getCreateAddressMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateAddress",
      requestType = CreateAddressReq.class,
      responseType = CreateAddressReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<CreateAddressReq,
      CreateAddressReply> getCreateAddressMethod() {
    io.grpc.MethodDescriptor<CreateAddressReq, CreateAddressReply> getCreateAddressMethod;
    if ((getCreateAddressMethod = UserGrpc.getCreateAddressMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getCreateAddressMethod = UserGrpc.getCreateAddressMethod) == null) {
          UserGrpc.getCreateAddressMethod = getCreateAddressMethod =
              io.grpc.MethodDescriptor.<CreateAddressReq, CreateAddressReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateAddress"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  CreateAddressReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  CreateAddressReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("CreateAddress"))
              .build();
        }
      }
    }
    return getCreateAddressMethod;
  }

  private static volatile io.grpc.MethodDescriptor<GetAddressReq,
      GetAddressReply> getGetUserAddressMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserAddress",
      requestType = GetAddressReq.class,
      responseType = GetAddressReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<GetAddressReq,
      GetAddressReply> getGetUserAddressMethod() {
    io.grpc.MethodDescriptor<GetAddressReq, GetAddressReply> getGetUserAddressMethod;
    if ((getGetUserAddressMethod = UserGrpc.getGetUserAddressMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getGetUserAddressMethod = UserGrpc.getGetUserAddressMethod) == null) {
          UserGrpc.getGetUserAddressMethod = getGetUserAddressMethod =
              io.grpc.MethodDescriptor.<GetAddressReq, GetAddressReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserAddress"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  GetAddressReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  GetAddressReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("GetUserAddress"))
              .build();
        }
      }
    }
    return getGetUserAddressMethod;
  }

  private static volatile io.grpc.MethodDescriptor<CreateProfileReq,
      CreateProfileReply> getCreateProfileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateProfile",
      requestType = CreateProfileReq.class,
      responseType = CreateProfileReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<CreateProfileReq,
      CreateProfileReply> getCreateProfileMethod() {
    io.grpc.MethodDescriptor<CreateProfileReq, CreateProfileReply> getCreateProfileMethod;
    if ((getCreateProfileMethod = UserGrpc.getCreateProfileMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getCreateProfileMethod = UserGrpc.getCreateProfileMethod) == null) {
          UserGrpc.getCreateProfileMethod = getCreateProfileMethod =
              io.grpc.MethodDescriptor.<CreateProfileReq, CreateProfileReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateProfile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  CreateProfileReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  CreateProfileReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("CreateProfile"))
              .build();
        }
      }
    }
    return getCreateProfileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<GetProfileReq,
      GetProfileReply> getGetUserProfileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserProfile",
      requestType = GetProfileReq.class,
      responseType = GetProfileReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<GetProfileReq,
      GetProfileReply> getGetUserProfileMethod() {
    io.grpc.MethodDescriptor<GetProfileReq, GetProfileReply> getGetUserProfileMethod;
    if ((getGetUserProfileMethod = UserGrpc.getGetUserProfileMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getGetUserProfileMethod = UserGrpc.getGetUserProfileMethod) == null) {
          UserGrpc.getGetUserProfileMethod = getGetUserProfileMethod =
              io.grpc.MethodDescriptor.<GetProfileReq, GetProfileReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserProfile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  GetProfileReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  GetProfileReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("GetUserProfile"))
              .build();
        }
      }
    }
    return getGetUserProfileMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static UserStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserStub>() {
        @java.lang.Override
        public UserStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserStub(channel, callOptions);
        }
      };
    return UserStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static UserBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserBlockingStub>() {
        @java.lang.Override
        public UserBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserBlockingStub(channel, callOptions);
        }
      };
    return UserBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static UserFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserFutureStub>() {
        @java.lang.Override
        public UserFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserFutureStub(channel, callOptions);
        }
      };
    return UserFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void health(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<HealthReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getHealthMethod(), responseObserver);
    }

    /**
     */
    default void getUser(GetUserReq request,
        io.grpc.stub.StreamObserver<GetUserReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserMethod(), responseObserver);
    }

    /**
     */
    default void getUserByUsername(GetUserByUsernameReq request,
        io.grpc.stub.StreamObserver<GetUserByUsernameReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserByUsernameMethod(), responseObserver);
    }

    /**
     */
    default void resetUserPassword(ResetUserPasswordReq request,
        io.grpc.stub.StreamObserver<ResetUserPasswordReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getResetUserPasswordMethod(), responseObserver);
    }

    /**
     */
    default void createUser(CreateUserReq request,
        io.grpc.stub.StreamObserver<CreateUserReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateUserMethod(), responseObserver);
    }

    /**
     */
    default void listUser(ListUserReq request,
        io.grpc.stub.StreamObserver<ListUserReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListUserMethod(), responseObserver);
    }

    /**
     */
    default void verifyPassword(VerifyPasswordReq request,
        io.grpc.stub.StreamObserver<VerifyPasswordReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getVerifyPasswordMethod(), responseObserver);
    }

    /**
     */
    default void createAddress(CreateAddressReq request,
        io.grpc.stub.StreamObserver<CreateAddressReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateAddressMethod(), responseObserver);
    }

    /**
     */
    default void getUserAddress(GetAddressReq request,
        io.grpc.stub.StreamObserver<GetAddressReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserAddressMethod(), responseObserver);
    }

    /**
     */
    default void createProfile(CreateProfileReq request,
        io.grpc.stub.StreamObserver<CreateProfileReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateProfileMethod(), responseObserver);
    }

    /**
     */
    default void getUserProfile(GetProfileReq request,
        io.grpc.stub.StreamObserver<GetProfileReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserProfileMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service User.
   */
  public static abstract class UserImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return UserGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service User.
   */
  public static final class UserStub
      extends io.grpc.stub.AbstractAsyncStub<UserStub> {
    private UserStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserStub(channel, callOptions);
    }

    /**
     */
    public void health(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<HealthReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getHealthMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUser(GetUserReq request,
        io.grpc.stub.StreamObserver<GetUserReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserByUsername(GetUserByUsernameReq request,
        io.grpc.stub.StreamObserver<GetUserByUsernameReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserByUsernameMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void resetUserPassword(ResetUserPasswordReq request,
        io.grpc.stub.StreamObserver<ResetUserPasswordReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getResetUserPasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createUser(CreateUserReq request,
        io.grpc.stub.StreamObserver<CreateUserReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listUser(ListUserReq request,
        io.grpc.stub.StreamObserver<ListUserReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void verifyPassword(VerifyPasswordReq request,
        io.grpc.stub.StreamObserver<VerifyPasswordReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getVerifyPasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createAddress(CreateAddressReq request,
        io.grpc.stub.StreamObserver<CreateAddressReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateAddressMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserAddress(GetAddressReq request,
        io.grpc.stub.StreamObserver<GetAddressReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserAddressMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createProfile(CreateProfileReq request,
        io.grpc.stub.StreamObserver<CreateProfileReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateProfileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserProfile(GetProfileReq request,
        io.grpc.stub.StreamObserver<GetProfileReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserProfileMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service User.
   */
  public static final class UserBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<UserBlockingStub> {
    private UserBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserBlockingStub(channel, callOptions);
    }

    /**
     */
    public HealthReply health(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getHealthMethod(), getCallOptions(), request);
    }

    /**
     */
    public GetUserReply getUser(GetUserReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public GetUserByUsernameReply getUserByUsername(GetUserByUsernameReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserByUsernameMethod(), getCallOptions(), request);
    }

    /**
     */
    public ResetUserPasswordReply resetUserPassword(ResetUserPasswordReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getResetUserPasswordMethod(), getCallOptions(), request);
    }

    /**
     */
    public CreateUserReply createUser(CreateUserReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public ListUserReply listUser(ListUserReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public VerifyPasswordReply verifyPassword(VerifyPasswordReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getVerifyPasswordMethod(), getCallOptions(), request);
    }

    /**
     */
    public CreateAddressReply createAddress(CreateAddressReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateAddressMethod(), getCallOptions(), request);
    }

    /**
     */
    public GetAddressReply getUserAddress(GetAddressReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserAddressMethod(), getCallOptions(), request);
    }

    /**
     */
    public CreateProfileReply createProfile(CreateProfileReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateProfileMethod(), getCallOptions(), request);
    }

    /**
     */
    public GetProfileReply getUserProfile(GetProfileReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserProfileMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service User.
   */
  public static final class UserFutureStub
      extends io.grpc.stub.AbstractFutureStub<UserFutureStub> {
    private UserFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<HealthReply> health(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getHealthMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<GetUserReply> getUser(
        GetUserReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<GetUserByUsernameReply> getUserByUsername(
        GetUserByUsernameReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserByUsernameMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<ResetUserPasswordReply> resetUserPassword(
        ResetUserPasswordReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getResetUserPasswordMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<CreateUserReply> createUser(
        CreateUserReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<ListUserReply> listUser(
        ListUserReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<VerifyPasswordReply> verifyPassword(
        VerifyPasswordReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getVerifyPasswordMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<CreateAddressReply> createAddress(
        CreateAddressReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateAddressMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<GetAddressReply> getUserAddress(
        GetAddressReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserAddressMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<CreateProfileReply> createProfile(
        CreateProfileReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateProfileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<GetProfileReply> getUserProfile(
        GetProfileReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserProfileMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_HEALTH = 0;
  private static final int METHODID_GET_USER = 1;
  private static final int METHODID_GET_USER_BY_USERNAME = 2;
  private static final int METHODID_RESET_USER_PASSWORD = 3;
  private static final int METHODID_CREATE_USER = 4;
  private static final int METHODID_LIST_USER = 5;
  private static final int METHODID_VERIFY_PASSWORD = 6;
  private static final int METHODID_CREATE_ADDRESS = 7;
  private static final int METHODID_GET_USER_ADDRESS = 8;
  private static final int METHODID_CREATE_PROFILE = 9;
  private static final int METHODID_GET_USER_PROFILE = 10;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_HEALTH:
          serviceImpl.health((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<HealthReply>) responseObserver);
          break;
        case METHODID_GET_USER:
          serviceImpl.getUser((GetUserReq) request,
              (io.grpc.stub.StreamObserver<GetUserReply>) responseObserver);
          break;
        case METHODID_GET_USER_BY_USERNAME:
          serviceImpl.getUserByUsername((GetUserByUsernameReq) request,
              (io.grpc.stub.StreamObserver<GetUserByUsernameReply>) responseObserver);
          break;
        case METHODID_RESET_USER_PASSWORD:
          serviceImpl.resetUserPassword((ResetUserPasswordReq) request,
              (io.grpc.stub.StreamObserver<ResetUserPasswordReply>) responseObserver);
          break;
        case METHODID_CREATE_USER:
          serviceImpl.createUser((CreateUserReq) request,
              (io.grpc.stub.StreamObserver<CreateUserReply>) responseObserver);
          break;
        case METHODID_LIST_USER:
          serviceImpl.listUser((ListUserReq) request,
              (io.grpc.stub.StreamObserver<ListUserReply>) responseObserver);
          break;
        case METHODID_VERIFY_PASSWORD:
          serviceImpl.verifyPassword((VerifyPasswordReq) request,
              (io.grpc.stub.StreamObserver<VerifyPasswordReply>) responseObserver);
          break;
        case METHODID_CREATE_ADDRESS:
          serviceImpl.createAddress((CreateAddressReq) request,
              (io.grpc.stub.StreamObserver<CreateAddressReply>) responseObserver);
          break;
        case METHODID_GET_USER_ADDRESS:
          serviceImpl.getUserAddress((GetAddressReq) request,
              (io.grpc.stub.StreamObserver<GetAddressReply>) responseObserver);
          break;
        case METHODID_CREATE_PROFILE:
          serviceImpl.createProfile((CreateProfileReq) request,
              (io.grpc.stub.StreamObserver<CreateProfileReply>) responseObserver);
          break;
        case METHODID_GET_USER_PROFILE:
          serviceImpl.getUserProfile((GetProfileReq) request,
              (io.grpc.stub.StreamObserver<GetProfileReply>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getHealthMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.google.protobuf.Empty,
              HealthReply>(
                service, METHODID_HEALTH)))
        .addMethod(
          getGetUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              GetUserReq,
              GetUserReply>(
                service, METHODID_GET_USER)))
        .addMethod(
          getGetUserByUsernameMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              GetUserByUsernameReq,
              GetUserByUsernameReply>(
                service, METHODID_GET_USER_BY_USERNAME)))
        .addMethod(
          getResetUserPasswordMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              ResetUserPasswordReq,
              ResetUserPasswordReply>(
                service, METHODID_RESET_USER_PASSWORD)))
        .addMethod(
          getCreateUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              CreateUserReq,
              CreateUserReply>(
                service, METHODID_CREATE_USER)))
        .addMethod(
          getListUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              ListUserReq,
              ListUserReply>(
                service, METHODID_LIST_USER)))
        .addMethod(
          getVerifyPasswordMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              VerifyPasswordReq,
              VerifyPasswordReply>(
                service, METHODID_VERIFY_PASSWORD)))
        .addMethod(
          getCreateAddressMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              CreateAddressReq,
              CreateAddressReply>(
                service, METHODID_CREATE_ADDRESS)))
        .addMethod(
          getGetUserAddressMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              GetAddressReq,
              GetAddressReply>(
                service, METHODID_GET_USER_ADDRESS)))
        .addMethod(
          getCreateProfileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              CreateProfileReq,
              CreateProfileReply>(
                service, METHODID_CREATE_PROFILE)))
        .addMethod(
          getGetUserProfileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              GetProfileReq,
              GetProfileReply>(
                service, METHODID_GET_USER_PROFILE)))
        .build();
  }

  private static abstract class UserBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    UserBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return GRPCUserServiceClientTest.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("User");
    }
  }

  private static final class UserFileDescriptorSupplier
      extends UserBaseDescriptorSupplier {
    UserFileDescriptorSupplier() {}
  }

  private static final class UserMethodDescriptorSupplier
      extends UserBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    UserMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (UserGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new UserFileDescriptorSupplier())
              .addMethod(getHealthMethod())
              .addMethod(getGetUserMethod())
              .addMethod(getGetUserByUsernameMethod())
              .addMethod(getResetUserPasswordMethod())
              .addMethod(getCreateUserMethod())
              .addMethod(getListUserMethod())
              .addMethod(getVerifyPasswordMethod())
              .addMethod(getCreateAddressMethod())
              .addMethod(getGetUserAddressMethod())
              .addMethod(getCreateProfileMethod())
              .addMethod(getGetUserProfileMethod())
              .build();
        }
      }
    }
    return result;
  }
}

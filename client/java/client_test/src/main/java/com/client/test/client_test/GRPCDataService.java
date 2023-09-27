package com.client.test.client_test;

import com.client.test.pb.grpc.UserGrpc;

import io.grpc.ManagedChannel;
import io.grpc.netty.NettyChannelBuilder;

public class GRPCDataService {
	private UserGrpc.UserBlockingStub userBlockingStub;
	private UserGrpc.UserStub userStub;
	private UserGrpc.UserFutureStub userFutureStub;
	
	private String baseUrl="localhost";
	private int port=7779;
	private ManagedChannel channel;
	
	public GRPCDataService() {
		userBlockingStub=UserGrpc.newBlockingStub(getChannel());
		userStub=UserGrpc.newStub(getChannel());
		userFutureStub=UserGrpc.newFutureStub(getChannel());
	}
	
	
	public UserGrpc.UserBlockingStub getUserBlockingStub() {
		return userBlockingStub;
	}


	public UserGrpc.UserStub getUserStub() {
		return userStub;
	}



	public UserGrpc.UserFutureStub getUserFutureStub() {
		return userFutureStub;
	}


	public ManagedChannel getChannel() {
		if ( channel==null ||  channel.isShutdown() || channel.isTerminated()){
			channel = NettyChannelBuilder.forAddress(baseUrl, port).usePlaintext().build();
		}
		return channel;
	}

}

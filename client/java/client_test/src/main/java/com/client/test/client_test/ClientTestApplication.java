package com.client.test.client_test;

import com.client.test.pb.grpc.HealthReply;
import com.client.test.pb.grpc.UserGrpc;
import com.google.protobuf.Empty;

public class ClientTestApplication {
	
	private GRPCDataService grpcService=new GRPCDataService();

	public static void main(String[] args) {
		new ClientTestApplication().clientBlockingTest();
	}

	public void clientBlockingTest() {
		UserGrpc.UserBlockingStub blockingStub=grpcService.getUserBlockingStub();
		HealthReply reply=blockingStub.health(Empty.getDefaultInstance());
		
		System.out.println(reply.getOk());
	}
}

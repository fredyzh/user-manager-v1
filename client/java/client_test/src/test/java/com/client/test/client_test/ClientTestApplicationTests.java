package com.client.test.client_test;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;

import org.apache.logging.log4j.Level;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.junit.jupiter.api.Test;

import com.client.test.pb.grpc.HealthReply;
import com.client.test.pb.grpc.ListUserReply;
import com.client.test.pb.grpc.ListUserReq;
import com.client.test.pb.grpc.User;
import com.client.test.pb.grpc.UserGrpc;
import com.google.protobuf.Empty;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;

class ClientTestApplicationTests {
	private Logger logger = LogManager.getLogger();
	private GRPCDataService grpcService=new GRPCDataService();

//	@Test
	public void testClientBlockingTest() {
		UserGrpc.UserBlockingStub blockingStub=grpcService.getUserBlockingStub();
		HealthReply reply=blockingStub.health(Empty.getDefaultInstance());
		
		assert(reply.getOk());
	}
	
	@Test
	public void testClientList() {
		UserGrpc.UserStub userStub=grpcService.getUserStub();
		
		final CountDownLatch finishLatch = new CountDownLatch(1);
		AtomicBoolean isError=new AtomicBoolean(false);
		
		List <User> users=new ArrayList<>();

		StreamObserver<ListUserReply> responseObserver = new StreamObserver<ListUserReply>() {

			@Override
			public void onNext(ListUserReply value) {
				logger.log(Level.INFO, () -> "Response: ");

				if(value.getResultsCount()>0) { 
					users.addAll(value.getResultsList()); 
				}
			}

			@Override
			public void onError(Throwable t) {
				Status status = Status.fromThrowable(t);
				logger.log(Level.INFO, () -> "save status: " + status);
				isError.getAndSet(true);
				finishLatch.countDown();
			}

			@Override
			public void onCompleted() {
				logger.log(Level.INFO, () -> "Finished RecordRoute");
				finishLatch.countDown();
			}

		};
		
		userStub.listUser(ListUserReq.getDefaultInstance(), responseObserver);
		ListUserReply reply=ListUserReply.getDefaultInstance();
		
		try {
				responseObserver.onNext(reply);
		} catch (RuntimeException e) {
			responseObserver.onError(e);
			throw e;
		}

		try {
			finishLatch.await(10, TimeUnit.SECONDS);
			
			logger.log(Level.INFO, () -> "Finished");
			
			if (isError.compareAndSet(true, false)) {
				throw new RuntimeException("failed retrieved data.");
			}
		} catch (InterruptedException e) {
		} 
		
		assert(users.size()>=1);
		
		UserGrpc.UserBlockingStub blockingStub=grpcService.getUserBlockingStub();
		ListUserReply reply1 = blockingStub.listUser(ListUserReq.getDefaultInstance());
		assert(reply1.getResultsCount()>=1);
	}
}

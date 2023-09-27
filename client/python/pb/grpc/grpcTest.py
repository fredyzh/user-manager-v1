import service_pb2_grpc
import service_pb2
import empty_pb2
import grpc


def run():
    print("run_test")
    with grpc.insecure_channel('localhost:7779') as channel:
        stub = service_pb2_grpc.UserStub(channel)   
        re=stub.GetUser(service_pb2.GetUserReq(id="1234"))
        print(re.id)
        
        rpl=stub.Health(empty_pb2.Empty())
        print(rpl)
        
        rul=stub.ListUser(service_pb2.ListUserReq())
        print(rul)
    print("test1")

if __name__ == "__main__":
    run()
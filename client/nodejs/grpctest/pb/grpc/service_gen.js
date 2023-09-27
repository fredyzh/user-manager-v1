const grpc = require("@grpc/grpc-js");
var protoLoader = require("@grpc/proto-loader");
const PROTO_PATH = "../service.proto";

const options = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};

var packageDefinition = protoLoader.loadSync(PROTO_PATH, options);

var protoDescriptor = grpc.loadPackageDefinition(packageDefinition).user.v1.service;

var service_user=protoDescriptor.User;

const client = new service_user('localhost:7779',
                                      grpc.credentials.createInsecure());

client.health({}, (error, reply) => {
    if (error!==null) {
        console.log(error);
    }else{
        console.log(reply);
    }
  });

client.listUser({},(error, reply) => {
    if (error!==null) {
        console.log(error);
    }else{
        console.log(reply);
    }
  });

client.getUser({id:"2345"},(error, reply) => {
if (error!==null) {
    console.log(error);
}else{
    console.log(reply);
}
});

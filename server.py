import logging
from concurrent import futures

import grpc
import api_pb2_grpc, api_pb2
from api_pb2_grpc import DNSProcessorServicer

from process import predict

class Service(DNSProcessorServicer):
    def IsEval(self, request, context):
        return api_pb2.Response(res=predict(request.domain))


def serve():
    print('start grpc server====>')
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    api_pb2_grpc.add_DNSProcessorServicer_to_server(Service(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
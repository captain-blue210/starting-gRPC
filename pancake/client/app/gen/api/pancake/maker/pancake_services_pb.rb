# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: pancake.proto for package 'pancake.maker'

require 'grpc'
require 'pancake_pb'

module Pancake
  module Maker
    module PancakeBakerService
      class Service

        include GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'pancake.maker.PancakeBakerService'

        # 指定されたメニューのパンを焼くメソッド
        rpc :Bake, ::Pancake::Maker::BakeRequest, ::Pancake::Maker::BakeResponse
        # メニューごとに焼いたパンケーキの数を返すメソッド
        rpc :Report, ::Pancake::Maker::ReportRequest, ::Pancake::Maker::ReportResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end

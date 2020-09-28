# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: image_upload.proto for package 'image.upload'

require 'grpc'
require 'image_upload_pb'

module Image
  module Upload
    module ImageUploadService
      class Service

        include GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'image.upload.ImageUploadService'

        rpc :Upload, stream(::Image::Upload::ImageUploadRequest), ::Image::Upload::ImageUploadResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end

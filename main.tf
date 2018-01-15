provider "sampleclassis" {
  url = "http://localhost:3000"
  email = "test@test.com"
  password = "qwerpoiu"
}


resource "sampleclassis_aws_spot_group" "test" {
  group_name = "what"
  desired_qty = "1"
  region = "us-west-2"
  quantity = "2"
  active = false
  vpc_id= ""
  image_id ="ami-77c74517"
  subnet_id = ""
  key_name = "drone"
  iam_fleet_role = ""
  default_device_size ="30"
  security_groups = [""]
  instance_types = ["m3.medium"]
}
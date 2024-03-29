# Creating a zip of the lambda binary. This zip will be deployed to lambda.
# This assumes the lambda has been built as a binary called 'api' in a directory one level above this one.
data "archive_file" "lambda_zip" {
  type = "zip"

  source_file  = "${path.module}/../api"
  output_path = "${path.module}/api.zip"
}

data "aws_iam_policy_document" "lambda_trust_policy" {
  statement {
    actions    = ["sts:AssumeRole"]
    effect     = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

# Says the lambda 'pricipal' can assume a role.
resource "aws_iam_role" "lambda_role_http_lambda" {
  name               = "http_lambda_role"
  assume_role_policy = data.aws_iam_policy_document.lambda_trust_policy.json
}

# Says what you can actually do when you assume the role.
resource "aws_iam_role_policy_attachment" "lambda_policy_http_lambda" {
  role       = "${aws_iam_role.lambda_role_http_lambda.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "http_lambda" {
  filename         = data.archive_file.lambda_zip.output_path
  function_name    = "lambda-rest-api-demo"
  handler          = "api" # For golang, the handler function must match the name of the binary present when the archive is unziped.
  role             = aws_iam_role.lambda_role_http_lambda.arn
  runtime          = "go1.x"
  source_code_hash = data.archive_file.lambda_zip.output_base64sha256
}

resource "aws_cloudwatch_log_group" "http_lambda" {
  name = "/aws/lambda/${aws_lambda_function.http_lambda.function_name}"

  retention_in_days = 5
}

resource "aws_lambda_function_url" "function_url" {
  function_name      = aws_lambda_function.http_lambda.function_name
  authorization_type = "NONE"
}

output "function_url" {
 value = aws_lambda_function_url.function_url.function_url
}
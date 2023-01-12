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
resource "aws_iam_role" "lambda_role" {
  name               = "lambda_role"
  assume_role_policy = data.aws_iam_policy_document.lambda_trust_policy.json
}

# Says what you can actually do when you assume the role.
resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = "${aws_iam_role.lambda_role.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "http_lambda" {
  filename         = "api.zip"
  function_name    = "http-lambda"
  handler          = "api" # For golang, the handler function must match the name of the binary present when the archive is unziped.
  role             = aws_iam_role.lambda_role.arn
  runtime          = "go1.x"
  source_code_hash = filebase64sha256("api.zip")
}

# 대전대 학식 카카오봇 스킬서버 (Serverless)

해야할 일 :

-   [x] Dev / Staging / production 환경으로 인프라 나누기
-   [ ] CI/CD 구축 (Jenkins)
-   [ ] CI/CD 구축 후 DynamoDB 도입

참고 :

| 목록                                   | 링크                                                                                                                              |
| -------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| Go Lambda function                     | [Link](https://docs.aws.amazon.com/ko_kr/lambda/latest/dg/golang-handler.html)                                                    |
| AWS Lambda & API Gateway               | [Link](https://learn.hashicorp.com/tutorials/terraform/lambda-api-gateway)                                                        |
| AWS s3 bucket notofication             | [Link](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_notification)                        |
| AWS cloudwatch event rule              | [Link](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_rule)                         |
| terraform lambda example               | [Link](https://github.com/snsinfu/terraform-lambda-example)                                                                       |
| AWS role                               | [Link](https://www.notion.so/IAM-ec1e72d874b540448401d7523693f3bb)                                                                |
| kakao open builder                     | [Link](https://i.kakao.com/docs/getting-started-overview#%EC%98%A4%ED%94%88%EB%B9%8C%EB%8D%94-%EC%86%8C%EA%B0%9C)                 |
| AWS Go SDK                             | [Link](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/)                                                                 |
| getting-error-fork-exec-var-task-main  | [Link](https://stackoverflow.com/questions/58133166/getting-error-fork-exec-var-task-main-no-such-file-or-directory-while-execut) |
| go apigateway example                  | [Link](https://github.com/serverless/examples/blob/master/aws-golang-http-get-post/postFolder/postExample.go)                     |
| AWS dynamoDB                           | [Link](https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/HowItWorks.CoreComponents.html)                     |
| AWS doc sdk example                    | [Link](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/go/example_code)                                                 |
| Devops Workshop                        | [Link](https://devops-art-factory.gitbook.io/devops-workshop/)                                                                    |
| Jenkins Terraform Integration          | [Link](https://www.coachdevops.com/2021/07/jenkins-terraform-integration-how-do.html)                                             |
| Jenkins Getting Started with Pipelines | [Link](https://www.jenkins.io/pipeline/getting-started-pipelines/)                                                                |
| Git Workflow                           | [Link](https://blog.ull.im/engineering/2019/06/25/git-workflow-for-ci-cd.html)                                                    |

```bash
.
├── src
│   ├── crawling
│   │   ├── config
│   │   │   └── url.yaml
│   │   ├── cover.out
│   │   ├── crawling.go
│   │   └── crawling_test.go
│   ├── menu
│   │   ├── cover.out
│   │   ├── menu.go
│   │   └── menu_test.go
│   └── skill-server
│       ├── main.go
│       ├── request-from-bot.go
│       ├── response.go
│       ├── service.go
│       ├── skill_server.go
│       └── utils.go
└── terraform
    ├── apiGateway
    │   ├── _module
    │   │   └── apiGateway
    │   │       ├── apiGateway.tf
    │   │       └── variables.tf
    │   ├── dev
    │   ├── production
    │   ├── provider.tf
    │   └── staging
    ├── ec2
    │   ├── dh-prod
    │   └── provider.tf
    ├── iam
    │   ├── byeongmin
    │   └── dh-prod
    ├── init
    │   ├── byeongmin
    │   └── dh-prod
    ├── lambda
    │   ├── _module
    │   │   ├── crawling-lambda
    │   │   ├── menu-lambda
    │   │   └── skill-server-lambda
    │   ├── dev
    │   ├── outputs.tf
    │   ├── production
    │   ├── provider.tf
    │   └── staging
    ├── logs
    │   ├── _module
    │   ├── dev
    │   ├── production
    │   ├── provider.tf
    │   └── staging
    └── s3
        ├── _module
        ├── dev
        ├── outputs.tf
        ├── production
        ├── provider.tf
        └── staging
```

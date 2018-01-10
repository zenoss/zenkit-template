#!/bin/bash
#
# register-api.sh - register a REST API with Google Endpoints
#
# This script registers the REST API with Google Endpoints for a single
# GCP project. The command line arguments are the target project ID
# and the path to the swagger YAML defining the API.
#
# Example: 
#    $ ./register-api.sh zing-testing swagger/swagger.yaml
#    $ ./register-api.sh zing-staging-191519 swagger/swagger.yaml
#
# Notes:
#
# 1. This script must be run separately for each GCP project (test, 
#    staging, etc). 
#
# 2. At end of a successful registration, you should a line like this 
#    near the end of the output:
#    Service Configuration [2018-01-10r0] uploaded for service [graphql.endpoints.zing-staging-191519.cloud.goog]
#
#    The value in "[]" is the service configuration ID ("2018-01-10r0" 
#    in this example). This value must be added to the file(s)
#    env/deployment/<project>/<service>.env in the zing-deploy 
#    github repo.
#
# 3. GCP project names and IDs can be different. This script requires
#    the project ID; not the name.  You can use "gcloud projects list"
#    to see a list of projects you have access to.
# 
# 4. You must be logged into GCP with an account that has write access
#    to the Endpoints API.  
#      - You can use 'gcloud auth list' to see the current, active 
#        account for your shell.
#      - You can use 'gcloud auth login' to login with your Gsuite
#        credentials.
# 

function print_usage() {
	echo "USAGE: $0 PROJECT_ID SWAGGER_FILE"
}

if [ $# -ne 2 ]; then
	echo "ERROR: argument count ($#) incorrect"
	print_usage
	exit 1
fi

set -x
set -e

CLOUD_PROJECT=$1
SWAGGER_FILE=$2
TMP_FILE=/tmp/api-register.tmp.yaml

rm -f $TMP_FILE
gcloud --project ${CLOUD_PROJECT} container clusters get-credentials cluster-1 --zone us-central1-a
sed "s/\${CLOUD_PROJECT}/${CLOUD_PROJECT}/g;" $SWAGGER_FILE >${TMP_FILE}
gcloud --project ${CLOUD_PROJECT} endpoints services deploy ${TMP_FILE}

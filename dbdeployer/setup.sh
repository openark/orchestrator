set -e

if [ -z "$MYSQL_USER" ]; then
    echo "Please provide MYSQL_USER as an environment variable when calling this script."
    exit 1
fi
if [ -z "$MYSQL_PASS" ]; then
    echo "Please provide MYSQL_PASS as an environment variable when calling this script."
    exit 1
fi
if [ -z "$MYSQL_HOST" ]; then
    echo "Please provide MYSQL_HOST as an environment variable when calling this script."
    exit 1
fi
if [ -z "$MYSQL_PORT" ]; then
    echo "Please provide MYSQL_PORT as an environment variable when calling this script."
    exit 1
fi
if [ -z "$MYSQL_VERSION" ]; then
    echo "Please provide MYSQL_VERSION as an environment variable when calling this script."
    exit 1
fi
if [ -z "$DESCRIPTION" ]; then
    echo "Please provide SANDBOX_PREFIX as an environment variable when calling this script."
    exit 1
fi


if [ -n "$CI" ]; then
  set -x
fi

if [ -n "$DEBUG" ]; then
    set -x
    export PS4='+(${BASH_SOURCE}:${LINENO}): ${FUNCNAME[0]:+${FUNCNAME[0]}(): }'
fi

if [ -z "$(uname | grep Darwin)" ]; then
  OS=linux
else
  OS=osx
fi

DBD_BASE_PATH="./dbdeployer"
DBD_SANDBOX_PATH="$DBD_BASE_PATH/sandboxes/$DESCRIPTION"
DBD_BIN_PATH="$DBD_BASE_PATH/bin"
DBD_VERSION="1.8.0"
DBD_DOWNLOAD_PATH="https://github.com/datacharmer/dbdeployer/releases/download/$DBD_VERSION"
DBD_DOWNLOAD_FILE="dbdeployer-$DBD_VERSION.$OS.tar.gz"
DBD_DOWNLOAD_URL="$DBD_DOWNLOAD_PATH/$DBD_DOWNLOAD_FILE"
DBD_CMD="$DBD_BIN_PATH/dbdeployer-$DBD_VERSION.$OS --sandbox-binary $DBD_BIN_PATH --sandbox-home $DBD_SANDBOX_PATH"

if [ "$OS" = "linux" ]; then
    MYSQL_FILE="$MYSQL_VERSION.tar.gz"
    MYSQL_DOWNLOAD_URL="https://raw.githubusercontent.com/datacharmer/mysql-docker-minimal/master/dbdata/$MYSQL_FILE"
else
    MYSQL_FILE="mysql-$MYSQL_VERSION-macos10.13-x86_64.tar.gz"
    MYSQL_DOWNLOAD_URL="https://dev.mysql.com/get/Downloads/MySQL-5.7/$MYSQL_FILE"
fi

echo "Creating working directories for dbdeployer..."
mkdir -p $DBD_SANDBOX_PATH
mkdir -p $DBD_BIN_PATH

echo "Checking if dbdeployer is installed..."
if ! [ -x "$(command -v $DBD_CMD)" ]; then
  echo "Not installed...starting install"
  rm -f $DBD_BIN_PATH/$DBD_DOWNLOAD_FILE*
  wget $DBD_DOWNLOAD_URL
  tar -vxzf $DBD_DOWNLOAD_FILE --directory $DBD_BIN_PATH
  rm $DBD_DOWNLOAD_FILE
else
  echo "Installation found! "
fi


echo "Checking if MySQL $MYSQL_VERSION is setup for dbdeployer..."
if [ -z "$($DBD_CMD available | grep $MYSQL_VERSION)" ]; then
  echo "Not found..."

  if [ ! -f $MYSQL_FILE ]; then
    echo "Downloading $MYSQL_FILE..."
    wget $MYSQL_DOWNLOAD_URL
  fi

  echo "Setting up MySQL $MYSQL_VERSION..."
  $DBD_CMD unpack $MYSQL_FILE --verbosity 0
  rm $MYSQL_FILE
else
  echo "MySQL $MYSQL_VERSION setup in dbdeployer! "
fi

echo "Rebuilding MySQL instance..."
$DBD_CMD deploy single $MYSQL_VERSION \
--force \
--db-user $MYSQL_USER \
--db-password $MYSQL_PASS \
--port $MYSQL_PORT \
--bind-address $MYSQL_HOST \

echo "Reporting status..."
$DBD_CMD global status 

echo "This dbdeployer sandbox can be controlled from:"
echo $DBD_SANDBOX_PATH/msb_${MYSQL_VERSION//./_} # replace 5.7.32 with 5_7_32
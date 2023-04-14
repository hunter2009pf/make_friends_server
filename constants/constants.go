package constants

// linux线上环境
// const PHOTO_SAVED_DIR = "/root/golang_projects/images"
// const BASE_URL = "http://47.101.214.35:10086"

// windows测试环境
const PHOTO_SAVED_DIR = "E:/go_projects/make_friends_server/images"
const BASE_URL = "http://192.168.132.111:10086"

// const BASE_URL = "http://192.168.2.6:10086"

const DATABASE_NAME = "makefriends"

var DEFAULT_NAME_PREFIX = []string{"可爱的", "勤劳的", "勇敢的", "热烈的", "大方的", "诚实的", "真诚的", "温柔的"}

const SUCCESS = 0

const ERROR_PARAMETER_EMPTY = 1

const ERROR_USER_ALREADY_EXIST = 2

const ERROR_ACCOUNT_PASSWORD_NOT_MATCH = 3

const ERROR_CREATE_FILE_FAILED = 4

const ERROR_COPY_FILE_FAILED = 5

const ERROR_PHOTO_NAME_EMPTY = 6

const ERROR_PERSONAL_INFO_EMPTY = 7

const ERROR_UPDATE_INFO_FAILED = 8

const ERROR_RANDOM_FIND_FAILED = 9

const ERROR_DECODE_FILE_FAILED = 10

const ERROR_ENCODE_FILE_FAILED = 11

const ERROR_LOGIN_CODE_EMPTY = 12

const ERROR_DELETE_PHOTO_FAILED = 13

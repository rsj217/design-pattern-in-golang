class Logger:
    def record_log(self):
        print(f"{self.__class__.__name__}: record a log")


class DBLogger(Logger):
    def record_log(self):
        print(f"{self.__class__.__name__}: store a log to db")


def make_log(logger: Logger):
    logger.record_log()


def client():
    log = Logger()
    make_log(log)      #  输出：Logger: record a log
    dblog = DBLogger()
    make_log(dblog)    #  输出：DBLogger: store a log to db

if __name__ == '__main__':
    client()
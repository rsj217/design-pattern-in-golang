from abc import ABC, abstractmethod


class AbstractLogger(ABC):
    @abstractmethod
    def record_log(self):
        raise NotImplementedError


class Logger(AbstractLogger):
    def record_log(self):
        print(f"{self.__class__.__name__}: record a log")


class DBLogger(AbstractLogger):
    def record_log(self):
        print(f"{self.__class__.__name__}: record a log")
        print(f"{self.__class__.__name__}: store a log to db")


def make_log(logger: AbstractLogger):
    logger.record_log()


def client():
    logger = Logger()
    make_log(logger)
    dblog = DBLogger()
    make_log(dblog)


if __name__ == '__main__':
    client()

from abc import ABC
from typing import Any


class TemplateMethod:

    def get_identify(self) -> str:
        raise NotImplementedError

    def send(self, msg: str, identify: str) -> bool:
        raise NotImplementedError

    def save(self, info: Any) -> bool:
        raise NotImplementedError

    def make_notify(self, msg: str) -> bool:
        print("Make Notify")
        identify = self.get_identify()
        ans = self.send(msg, identify)
        self.save(ans)
        return ans


class Push(TemplateMethod, ABC):

    def __init__(self, uid: int):
        self._uid = uid

    def get_identify(self) -> str:
        return f"{self._uid}-ios-14-110100"

    def send(self, msg: str, identify: str) -> bool:
        print(f"{identify} {self.__class__.__name__} msg: {msg}")
        return True

    def save(self, info: Any) -> bool:
        print(f"{self.__class__.__name__} save {info} to db")


class SMS(TemplateMethod, ABC):

    def __init__(self, uid: int):
        self._uid = uid

    def get_identify(self) -> str:
        return f"18519191001-{self._uid}"

    def send(self, msg: str, identify: str) -> bool:
        print(f"{identify} {self.__class__.__name__} msg: {msg}")
        return True

    def save(self, info: Any) -> bool:
        print(f"{self.__class__.__name__} record {info}  on log")


def main():
    uid = 110110
    msg = "hello world"

    push = Push(uid)
    push.make_notify(msg)


if __name__ == '__main__':
    main()

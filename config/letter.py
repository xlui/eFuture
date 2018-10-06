# Object to store future letters
import json


class Letter:
    def __init__(self, subject, receiver, receiveDate, content, sendDate) -> None:
        self.subject = subject
        self.receiver = receiver
        self.receiveDate = receiveDate
        self.content = content
        self.sendDate = sendDate
        super().__init__()

    def toJSON(self):
        return json.dumps(self.__toDict())

    def __toDict(self):
        return {
            'subject': self.subject,
            'receiver': self.receiver,
            'receiveDate': self.receiveDate,
            'sendDate': self.sendDate,
            'content': self.content
        }

    def __repr__(self):
        return self.__toDict().__repr__()

    def __str__(self):
        return self.__repr__()

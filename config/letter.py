# Object to store future letters
class Letter:
    def __init__(self, subject, sender, receiveDate, content, sendDate) -> None:
        self.subject = subject
        self.sender = sender
        self.receiveDate = receiveDate
        self.content = content
        self.sendDate = sendDate
        super().__init__()

    def __repr__(self):
        return {
            'subject': self.subject,
            'sender': self.sender,
            'receiveDate': self.receiveDate,
            'sendDate': self.sendDate,
            'content': self.content
        }.__repr__()

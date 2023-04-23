from PyQt6 import QtWidgets


class StartWindow(QtWidgets.QMainWindow):
    def __init__(self):
        super().__init__()

        self.setWindowTitle("Antivirus")

        layout = QtWidgets.QVBoxLayout()
        label = QtWidgets.QLabel("Сервис не запущен. Запустить?")
        button = QtWidgets.QPushButton("Запуск")

        layout.addWidget(label)
        layout.addWidget(button)

        widget = QtWidgets.QWidget()
        widget.setLayout(layout)

        self.setCentralWidget(widget)


app = QtWidgets.QApplication([])
window = StartWindow()
window.show()

app.exec()

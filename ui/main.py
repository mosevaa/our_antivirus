#здесь создаем граф интерфейс
from PyQt6 import QtWidgets
from PyQt6.QtWidgets import QApplication, QMainWindow


class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()

        self.setWindowTitle("Antivirus")
        file_path_line = QtWidgets.QLineEdit("Путь к файлу или директории")
        browse_button = QtWidgets.QPushButton("Browse...")
        status_label = QtWidgets.QLabel("Статус: что то")
        radio_buttons = QtWidgets.QButtonGroup()
        choice_status1 = QtWidgets.QRadioButton("Мониторинг")
        radio_buttons.addButton(choice_status1)
        choice_status2 = QtWidgets.QRadioButton("Сканирование по расписанию")
        radio_buttons.addButton(choice_status2)
        choice_status3 = QtWidgets.QRadioButton("Просканировать один раз")
        radio_buttons.addButton(choice_status3)
        scan_label = QtWidgets.QLabel("sds")
        scan_label.setStyleSheet("border: 1px solid black;")
        result_label = QtWidgets.QLabel("Результат: \n Что то \nЧто то еще \nЧто то еще")
        carantine_button = QtWidgets.QPushButton("Карантин")
        stop_service_button = QtWidgets.QPushButton("Остановить сервис")

        first_layout = QtWidgets.QHBoxLayout()
        first_layout.addWidget(file_path_line)
        first_layout.addWidget(browse_button)

        second_layout = QtWidgets.QHBoxLayout()
        second_layout.addWidget(status_label)
        second_layout.addWidget(choice_status1)
        second_layout.addWidget(choice_status2)
        second_layout.addWidget(choice_status3)

        third_layout = QtWidgets.QVBoxLayout()
        third_layout.addWidget(scan_label)

        buttons_layout = QtWidgets.QVBoxLayout()
        buttons_layout.addWidget(carantine_button)
        buttons_layout.addWidget(stop_service_button)

        fouth_layout = QtWidgets.QHBoxLayout()
        fouth_layout.addWidget(result_label)
        fouth_layout.addLayout(buttons_layout)

        main_layout = QtWidgets.QVBoxLayout()
        main_layout.addLayout(first_layout)
        main_layout.addLayout(second_layout)
        main_layout.addLayout(third_layout)
        main_layout.addLayout(fouth_layout)

        main_widget = QtWidgets.QWidget()
        main_widget.setLayout(main_layout)
        self.setCentralWidget(main_widget)


app = QApplication([])

window = MainWindow()
window.show()

app.exec()

import subprocess
import sys, os, time, threading
import main_window
import quarantine_delete_window
from PyQt5 import QtCore, QtGui, QtWidgets
from PyQt5.QtWidgets import QFileDialog
from subprocess import Popen, PIPE

os.environ['PYTHONIOENCODING'] = 'utf-8'


class MainWin(QtWidgets.QMainWindow):
    def __init__(self):
        QtWidgets.QWidget.__init__(self)
        self.ui = main_window.Ui_MainWindow()
        self.ui.setupUi(self)
        self.ui.pushButton_2.clicked.connect(self.start)
        self.ui.pushButton_5.clicked.connect(self.cleat_log_text_edit)
        self.ui.pushButton_3.clicked.connect(self.change_quarantine)

    def message(self, m):
        self.ui.textEdit_2.appendPlainText(m)

    def start(self):
        self.message('start')
        cmd = 'go run C:\\Users\\Lesya\\GolandProjects\\our_antivirus\\av\\main.go'
        p = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        stdout, stderr = p.communicate()
        self.ui.textEdit_2.appendPlainText(str(stdout.decode() + stderr.decode()))

    def cleat_log_text_edit(self):
        self.ui.textEdit_2.setPlainText('')

    def change_quarantine(self):
        pass


class QuarantineWin(QtWidgets.QMainWindow):
    def __init__(self, path):
        QtWidgets.QWidget.__init__(self)
        self.ui = quarantine_delete_window.Ui_MainWindow()
        self.ui.setupUi(self)
        self.scanning_path = path
        self.quarantine_path = ''  # path to quarantine folder
        self.quarantine_files = self.get_quarantine_files()

    def get_quarantine_files(self):
        return os.listdir(self.quarantine_path)


if __name__ == "__main__":
    app = QtWidgets.QApplication(sys.argv)
    MyApp = MainWin()
    MyApp.show()
    sys.exit(app.exec_())

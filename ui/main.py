import subprocess
import sys, os, time, threading
import main_window
import quarantine_delete_window
import set_time_window
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
        self.ui.pushButton.clicked.connect(self.getDirectory)

    def message(self, m):
        self.ui.textEdit_2.appendPlainText(m)

    def start(self):
        if self.ui.radioButton.isChecked():
            self.monitor()
        elif self.ui.radioButton_2.isChecked():
            self.schedule_scan()

    def cleat_log_text_edit(self):
        self.ui.textEdit_2.setPlainText('')

    def change_quarantine(self):
        pass

    @QtCore.pyqtSlot()
    def getDirectory(self):
        def getOpenFilesAndDirs(parent=None, caption='', directory='',
                        filter='', initialFilter='', options=None):
            def updateText():
                # update the contents of the line edit widget with the selected files
                selected = []
                for index in view.selectionModel().selectedRows():
                    selected.append('"{}"'.format(index.data()))
                self.ui.textEdit.setText(' '.join(selected))

            dialog = QtWidgets.QFileDialog(parent, windowTitle=caption)
            dialog.setFileMode(dialog.ExistingFiles)
            if options:
                dialog.setOptions(options)
            dialog.setOption(dialog.DontUseNativeDialog, True)
            if directory:
                dialog.setDirectory(directory)
            if filter: 
                dialog.setNameFilter(filter)
                if initialFilter:
                    dialog.selectNameFilter(initialFilter)

            # by default, if a directory is opened in file listing mode, 
            # QFileDialog.accept() shows the contents of that directory, but we 
            # need to be able to "open" directories as we can do with files, so we 
            # just override accept() with the default QDialog implementation which 
            # will just return exec_()
            dialog.accept = lambda: QtWidgets.QDialog.accept(dialog)

            # there are many item views in a non-native dialog, but the ones displaying 
            # the actual contents are created inside a QStackedWidget; they are a 
            # QTreeView and a QListView, and the tree is only used when the 
            # viewMode is set to QFileDialog.Details, which is not this case
            stackedWidget = dialog.findChild(QtWidgets.QStackedWidget)
            view = stackedWidget.findChild(QtWidgets.QListView)
            view.selectionModel().selectionChanged.connect(updateText)
            lineEdit = dialog.findChild(QtWidgets.QLineEdit)
            # clear the line edit contents whenever the current directory changes
            dialog.directoryEntered.connect(lambda: self.ui.textEdit.setText(''))

            dialog.exec_()
            return dialog.selectedFiles()
        fname = getOpenFilesAndDirs(self, "Open file or directory", "", "")[0]
        
        if fname:
            self.ui.textEdit.setText(fname)

    def monitor(self):
        pass

    def schedule_scan(self):
        self.message('start')
        cmd = 'go run C:\\Users\\Lesya\\GolandProjects\\our_antivirus\\av\\main.go'
        p = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        stdout, stderr = p.communicate()
        self.ui.textEdit_2.appendPlainText(str(stdout.decode() + stderr.decode()))


class SchedWin(QtWidgets.QMainWindow):
    def __init(self):
        QtWidgets.QWidget.__init__(self)
        self.ui = set_time_window.Ui_MainWindow()
        self.ui.setupUi(self)



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
    sched_win = SchedWin()
    MyApp.show()
    sys.exit(app.exec_())

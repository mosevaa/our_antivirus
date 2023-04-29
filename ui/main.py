import subprocess
import sys, os, time
from PyQt5 import QtCore, QtGui, QtWidgets
from subprocess import Popen, PIPE
from datetime import timedelta, datetime
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler

import main_window
import quarantine_delete_window

os.environ['PYTHONIOENCODING'] = 'utf-8'


class SchedScanner(QtCore.QThread):
    stdout = QtCore.pyqtSignal(str)
    next_scan = QtCore.pyqtSignal(str)

    def __init__(self, path, period):
        super().__init__()
        self.path = path
        self.period = period

    def run(self):
        while True:
            time.sleep(self.period)
            cmd = f"go run ../av/main.go -cmd scan_all_{self.path}"
            print(cmd)
            with Popen(cmd, stdout=PIPE) as p:
                while True:
                    text = p.stdout.read().decode("utf-8")
                    time.sleep(0.1)

                    if text == "":
                        break

                    self.stdout.emit(text)
                scan_time = datetime.now() + timedelta(seconds=self.period)
                self.next_scan.emit(str(scan_time.hour)+":"+str(scan_time.minute))


class Monitoring:
    def __init__(self):
        self.last_trigger = time.time()
        self.path = None


class MonHandler(FileSystemEventHandler):
    def on_created(self, event):
        if not os.path.isdir(event.src_path) and (time.time() - monitoring.last_trigger) > 1:
            monitoring.last_trigger = time.time()
            path = event.src_path.replace('\\', '/')
            print(path)
            my_app.message(
                f"Change in directory. Starting directory scan..."
            )
            print(111)
            cmd = f"go run ../av/main.go -cmd scan_all_{path}" #command to start scan
            print(1)
            p = subprocess.Popen(cmd, shell=True, stdout=PIPE)
            print(2)
            stdout, stderr = p.communicate()
            my_app.message(str(stdout.decode('utf8', errors='replace')+stderr.decode('utf8', errors='replace')))

    def on_modified(self, event):
        if not os.path.isdir(event.src_path) and (time.time() - monitoring.last_trigger) > 10000000:
            monitoring.last_trigger = time.time()
            my_app.message(
                f"Change in directory. Starting directory scan..."
            )
            print('here')
            cmd = f"go run ../av/main.go -cmd scan_all_{event.src_path}" #command to start scan
            p = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            stdout, stderr = p.communicate()
            my_app.message(str(stdout.decode('utf8')+stderr.decode('utf8', errors='replace')))


class MainWin(QtWidgets.QMainWindow):
    def __init__(self):
        QtWidgets.QWidget.__init__(self)
        self.ui = main_window.Ui_MainWindow()
        self.ui.setupUi(self)
        self.ui.pushButton_2.clicked.connect(self.start)
        self.ui.pushButton_4.clicked.connect(self.cleat_log_text_edit)
        self.ui.pushButton_3.clicked.connect(self.change_quarantine)
        self.ui.pushButton.clicked.connect(self.get_directory)
        self.ui.pushButton_5.clicked.connect(self.stop_server)

    def message(self, m):
        self.ui.textEdit_2.appendPlainText(m)

    def start(self):
        if self.ui.radioButton.isChecked():
            if self.ui.textEdit.toPlainText() != "":
                monitoring.path = self.ui.textEdit.toPlainText()
                self.message(f"Set path to monitor: {monitoring.path}")
                self.event_handler = MonHandler()
                self.observer = Observer()
                self.observer.schedule(self.event_handler, path=monitoring.path, recursive=True)
                self.observer.start()
            else:
                self.message("Please, browse file or directory to scan")
        elif self.ui.radioButton_2.isChecked():
            period, ok = QtWidgets.QInputDialog.getInt(self, 'Int', 'Enter')
            if ok:
                self.message(f"Period to scan: {period}")
                if self.ui.textEdit.toPlainText() != "" :
                    self.sched_scanner = SchedScanner(
                        self.ui.textEdit.toPlainText(),
                        period
                    )
                    self.sched_scanner.stdout.connect(self.ui.textEdit_2.appendPlainText)
                    self.sched_scanner.next_scan.connect(self.ui.textEdit_2.appendPlainText)
                    self.sched_scanner.start()
                else:
                    self.message("Please, browse file or directory to scan")

    def cleat_log_text_edit(self):
        self.ui.textEdit_2.setPlainText('')

    def change_quarantine(self):
        cmd = f"go run ../av/main.go -cmd quarantine" #command to start scan
        p = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        stdout, stderr = p.communicate()
        self.message(str(stdout.decode('utf8')+stderr.decode('utf8', errors='replace')))

    def stop_server(self):
        if self.ui.radioButton.isChecked():
            self.observer.stop()
            self.message("Stop monitoring directory")
        elif self.ui.radioButton_2.isChecked():
            self.sched_scanner.terminate()
            self.message("Stop scheduler scanning")

    @QtCore.pyqtSlot()
    def get_directory(self):
        def get_open_files_and_dirs(parent=None, caption='', directory='',
                        filter='', initialFilter='', options=None):
            def update_text():
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
            stacked_widget = dialog.findChild(QtWidgets.QStackedWidget)
            view = stacked_widget.findChild(QtWidgets.QListView)
            view.selectionModel().selectionChanged.connect(update_text)
            line_edit = dialog.findChild(QtWidgets.QLineEdit)
            # clear the line edit contents whenever the current directory changes
            dialog.directoryEntered.connect(lambda: self.ui.textEdit.setText(''))

            dialog.exec_()
            return dialog.selectedFiles()
        fname = get_open_files_and_dirs(self, "Open file or directory", "", "")[0]
        
        if fname:
            self.ui.textEdit.setText(fname)


class QuarantineWin(QtWidgets.QMainWindow):
    def __init__(self, path, parent=None):
        QtWidgets.QWidget.__init__(self, parent)
        self.ui = quarantine_delete_window.Ui_MainWindow()
        self.ui.setupUi(self)
        self.scanning_path = path
        self.quarantine_path = ''  # path to quarantine folder
        self.quarantine_files = self.get_quarantine_files()

    def get_quarantine_files(self):
        return os.listdir(self.quarantine_path)


if __name__ == "__main__":
    app = QtWidgets.QApplication(sys.argv)
    my_app = MainWin()
    monitoring = Monitoring()
    my_app.show()
    sys.exit(app.exec_())

"""
Process voters for music competitions
"""
from PySide6.QtWidgets import QApplication, QMainWindow, QTableWidget, QTableWidgetItem, QLineEdit, QGridLayout, \
    QHBoxLayout, QVBoxLayout, QGroupBox, QWidget, QLabel, QPlainTextEdit, QPushButton, QToolBar, QToolButton
import sys


class Ida(QMainWindow):
    def __init__(self):
        super().__init__()
        self.title = "Ida"

        self._init_window()
        self._init_ui()

        self.adjustSize()
        self.setMinimumSize(self.width(), self.height())
        self.show()

    def _init_window(self):
        self.setWindowTitle(self.title)

    def _init_ui(self):
        self.voter_name_le = QLineEdit()
        self.vote_details_te = QPlainTextEdit()
        self.parsed_votes_te = QPlainTextEdit()
        self.parsed_votes_te.setReadOnly(True)

        self.load_contest_btn = QToolButton()
        self.load_contest_btn.setText("Load Contest")
        self.parse_votes_btn = QToolButton()
        self.parse_votes_btn.setText("Parse Votes")
        self.reset_voter_btn = QToolButton()
        self.reset_voter_btn.setText("Reset Voter Details")

        toolbar = QToolBar()
        toolbar.addWidget(self.load_contest_btn)
        toolbar.addWidget(self.parse_votes_btn)
        toolbar.addWidget(self.reset_voter_btn)

        voter_details_grid = QVBoxLayout()
        voter_details_grid.addWidget(self.voter_name_le)
        voter_details_grid.addWidget(self.vote_details_te)
        voter_details_group = QGroupBox("Voter Details")
        voter_details_group.setLayout(voter_details_grid)

        parsed_votes_grid = QVBoxLayout()
        parsed_votes_grid.addWidget(self.parsed_votes_te)
        parsed_votes_group = QGroupBox("Parsed Votes")
        parsed_votes_group.setLayout(parsed_votes_grid)

        content_layout = QHBoxLayout()
        content_layout.addWidget(voter_details_group)
        content_layout.addWidget(parsed_votes_group)

        layout = QVBoxLayout()
        layout.addWidget(toolbar)
        layout.addLayout(content_layout)

        central_widget = QWidget()
        central_widget.setLayout(layout)
        self.setCentralWidget(central_widget)


def main():
    app = QApplication(sys.argv)
    app.setStyle("Fusion")
    main_window = Ida()
    sys.exit(app.exec_())

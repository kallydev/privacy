import logging
import sqlite3
import threading
import time


class Scanner:

    def __init__(self, database_name, file_name):
        logging.basicConfig(
            level=logging.DEBUG,
            format="%(asctime)s %(levelname)s %(message)s",
            datefmt='%Y-%m-%d %H:%M:%S',
        )
        self.logger = logging.getLogger()
        self.database_connection = None
        self.database_name = database_name
        self.file_name = file_name
        self.file_rows = 0
        self.handle_total = 0
        self.handle_invalid = 0
        self.handle_queue = 0
        self.cancel_print_insertion_speed = None

    def connect_database(self):
        self.database_connection = sqlite3.connect(self.database_name)

    def close_database(self):
        self.database_connection.close()

    def insert_qq_and_phone(self, id, qq, phone):
        cursor = self.database_connection.cursor()
        try:
            cursor.execute("INSERT INTO qq VALUES (?, ?, ?);", (id, qq, phone))
        except sqlite3.IntegrityError:
            self.handle_invalid += 1
        finally:
            self.handle_total += 1
            self.handle_queue += 1

    def start_insertion_speed(self):
        event = threading.Event()

        def print_insertion_speed():
            handle_total = self.handle_total
            while not event.wait(1):
                if self.handle_total - handle_total == 0:
                    continue
                self.logger.info("{}/s, {}/{} progress, {} rows are invalid, {} seconds left".format(
                    self.handle_total - handle_total,
                    self.handle_total,
                    self.file_rows,
                    self.handle_invalid,
                    (self.file_rows - self.handle_total) / (self.handle_total - handle_total),
                ))
                handle_total = self.handle_total

        threading.Thread(target=print_insertion_speed).start()
        return event.set

    def start(self):
        # Get the number of file rows
        self.logger.info("start scanning file lines")
        start_time = time.time()
        with open(self.file_name) as file:
            self.file_rows = 0
            for _ in file:
                self.file_rows += 1
        end_time = time.time()
        self.logger.info("scan completed, there are a total of {} lines, and it taken {} seconds".format(
            self.file_rows,
            end_time - start_time,
        ))
        # Insert QQ and phone numbers
        self.connect_database()
        self.cancel_print_insertion_speed = self.start_insertion_speed()
        with open(self.file_name) as file:
            for line in file:
                line = line.strip()
                data = line.split("----")
                if len(data) < 2:
                    self.handle_invalid += 1
                    self.handle_total += 1
                    continue
                phone = data[-1]
                for qq in data[:-1]:
                    self.insert_qq_and_phone(self.handle_total, qq, phone)
                if self.handle_queue >= 400000:
                    self.database_connection.commit()
                    self.handle_queue = 0
        self.database_connection.commit()
        self.cancel_print_insertion_speed()
        self.close_database()
        self.logger.info("completed, insert {} rows, {} rows of invalid data".format(
            self.handle_total,
            self.handle_invalid,
        ))
        exit()


if __name__ == '__main__':
    scanner = Scanner("database/database.db", "source/6.9更新总库.txt")
    scanner.start()

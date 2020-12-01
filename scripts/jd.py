import logging
import sqlite3
import threading
import time


class Converter(object):
    def __init__(self, database_path, file_path):
        logging.basicConfig(
            level=logging.DEBUG,
            format="%(asctime)s %(levelname)s %(message)s",
            datefmt='%Y-%m-%d %H:%M:%S',
        )
        self.logger = logging.getLogger()
        self.database_connection = None
        self.database_path = database_path
        self.file_path = file_path
        self.file_rows = 0
        self.handle_total = 0
        self.handle_invalid = 0
        self.handle_queue = 0
        self.cancel_print_insertion_speed = None

    def connect_database(self):
        self.database_connection = sqlite3.connect(self.database_path)

    def close_database(self):
        self.database_connection.close()

    def insert(self, id, name, nickname, password, email, id_number, phone_number):
        cursor = self.database_connection.cursor()
        try:
            cursor.execute(
                "INSERT INTO jd VALUES (?, ?, ?, ?, ?, ?, ?);",
                (id, name, nickname, password, email, id_number, phone_number)
            )
        except sqlite3.IntegrityError:
            self.handle_invalid += 1
        finally:
            self.handle_total += 1
            self.handle_queue += 1
        pass

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
        with open(self.file_path) as file:
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
        with open(self.file_path) as file:
            for line in file:
                try:
                    dataset = line.strip().split("---")
                    name = dataset[0]
                    nickname = dataset[1]
                    password = dataset[2]
                    email = dataset[3]
                    id_number = dataset[4]
                    phone_number = dataset[5]
                except IndexError:
                    self.handle_invalid += 1
                    pass
                finally:
                    self.handle_total += 1
                self.insert(self.handle_total, name, nickname, password, email, id_number, phone_number)
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
    converter = Converter("database/database.db", "source/www_jd_com_12g.txt")
    converter.start()

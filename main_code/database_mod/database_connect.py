#!/usr/bin/env python 
# -*- coding: utf-8 -*- j
import pymysql
import datetime
import prettytable
import os

state_machine = {
    1:'spare',
    2:'insert',
    3:'change',
    4:'delete',
    5:'end',
    6:'search',
}

class Database_control(object):
    def __init__(self):
        self.db = pymysql.connect("121.199.40.243","YuanCheng_user","_Hexagram_user","test_python")
        self.cursor = self.db.cursor()
        
    def print_data(self):
        os.system('cls')
        print(state_machine)
        sql = "select * from events_arrangement"
        self.__direct_database_control(sql)
            
        self.cursor.execute(sql)
        results = self.cursor.fetchall()
        self.__direct_print(results)

    def change_database(self):
        colums = {
            1:'events_end_time',
            2:'events_title',
            3:'events_advance_waring_time',
            4:'extre_sth'
        }
        temp_id = int(input("please the id you want to change"))
        temp=int(input("""please what you want to change\n 
            (end_time:1, sth_title:2, advance_warning_time:3,extre_sth:4):"""))
        
        if temp == 1:
            temp_input = input("end_time:")
        elif temp == 2:
            temp_input = input("sth_title:")
        elif temp == 4:
            temp_input = input("extre_sth")
        else:
            # temp == 3
            temp_input = input("advance_warning_time:")
        sql = \
            """
            UPDATE events_arrangement
            SET %s='%s'
            WHERE id=%d  
            """ % (colums[temp],temp_input,temp_id)
        self.__direct_database_control(sql)
        
        return state_machine[1]

    def insert_into_database(self):
        title = input("please input sth")
        end_time = input("please input end time,\"- -\"")
        if end_time == '':
            end_time = '1111-11-11 11:11:11'

        advance_warning_time = input("please input advance warning tima,\": :\"")
        if advance_warning_time == '':
            advance_warning_time = '11:11:11'
            
        extre_sth = input("please inpot extra sth")
        create_time = datetime.datetime.now()
        create_time_str = datetime.datetime.strftime(create_time,'%Y-%m-%d %H:%M:%S')

        sql = """
            insert into events_arrangement 
            (events_title,events_create_date,events_end_time,
            events_advance_waring_time,extre_sth)
            values
            ('%s','%s','%s','%s','%s')
            """ % (title,create_time_str,end_time,advance_warning_time,extre_sth)
            #print(sql)
        self.__direct_database_control(sql)
        return state_machine[1]

    def delete_id_row(self):
        row_id = int(input("Which id do you want to delete"))
        sql = """
            DELETE FROM events_arrangement
            WHERE id = %d
            """ % (row_id)
        self.__direct_database_control(sql)
        return state_machine[1]

    def __direct_database_control(self,sql):
        try:
            self.cursor.execute(sql)
            self.db.commit()
        except:
            print("wrong!!!")
            self.db.rollback()

    def __direct_print(self,temp):
        table = prettytable.PrettyTable(["ID","事务内容", \
            "创建时间","截至日期","提前预警时间","额外附加"])
        for row in temp:
            _id = row[0]
            _title = row[1]
            _create_date = str(row[2])
            _end_time = str(row[3])
            _advance_warning_time = str(row[4])
            _extre_sth = row[5]            
            table.add_row( \
            [_id,_title,_create_date[0:-3], \
            _end_time[0:-3],_advance_warning_time[0:-3],_extre_sth] \
            )
        print(table)

    def end_connect(self):
        self.db.close()
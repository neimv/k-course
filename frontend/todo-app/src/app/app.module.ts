import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AddTodoComponent } from './components/add-todo/add-todo.component';
import { TodoAddComponent } from './components/todo-add/todo-add.component';
import { TodoDetailsComponent } from './components/todo-details/todo-details.component';
import { TodoListComponent } from './components/todo-list/todo-list.component';
import { TodoUpdateComponent } from './components/todo-update/todo-update.component';

@NgModule({
  declarations: [
    AppComponent,
    AddTodoComponent,
    TodoAddComponent,
    TodoDetailsComponent,
    TodoListComponent,
    TodoUpdateComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

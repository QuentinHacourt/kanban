import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { KanbanComponent } from './kanban/kanban.component';
import { StoryDetailComponent } from './story-detail/story-detail.component';
import { CreateStoryComponent } from './create-story/create-story.component';

@NgModule({
  declarations: [AppComponent, KanbanComponent, StoryDetailComponent, CreateStoryComponent],
  imports: [BrowserModule, AppRoutingModule, FormsModule, HttpClientModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}

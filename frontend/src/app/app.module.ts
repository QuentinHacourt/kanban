import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { KanbanComponent } from './kanban/kanban.component';
import { StoryDetailComponent } from './story-detail/story-detail.component';
import { CreateStoryComponent } from './create-story/create-story.component';
import { RegisterDeveloperComponent } from './register-developer/register-developer.component';
import { DeveloperDetailsComponent } from './developer-details/developer-details.component';
import { DevelopersComponent } from './developers/developers.component';

@NgModule({
  declarations: [AppComponent, KanbanComponent, StoryDetailComponent, CreateStoryComponent, RegisterDeveloperComponent, DeveloperDetailsComponent, DevelopersComponent],
  imports: [BrowserModule, AppRoutingModule, FormsModule, HttpClientModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}

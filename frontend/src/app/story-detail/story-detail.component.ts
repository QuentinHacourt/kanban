import { Component, OnInit } from '@angular/core';
import { Story } from '../story';
import { StoryService } from '../story.service';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-story-detail',
  templateUrl: './story-detail.component.html',
  styleUrls: ['./story-detail.component.css'],
})
export class StoryDetailComponent implements OnInit {
  story: Story | undefined;

  constructor(
    private route: ActivatedRoute,
    private storyService: StoryService,
    private location: Location
  ) {}

  ngOnInit(): void {
    this.getStory();
  }

  getStory(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));
    this.storyService.getStory(id).subscribe((story) => (this.story = story));
  }

  goBack(): void {
    this.location.back();
  }

  save(): void {
    if (this.story) {
      this.storyService.updateStory(this.story).subscribe();
    }
  }
}

import {Component, OnDestroy, OnInit} from '@angular/core';
import {MatDialog, Sort} from '@angular/material';
import {ActivatedRoute} from '@angular/router';
import {find} from 'lodash';
import {interval, Subscription} from 'rxjs';
import {retry} from 'rxjs/operators';
import {AppConfigService} from '../app-config.service';
import {ApiService, UserService} from '../core/services';
import {AddSshKeyModalComponent} from '../shared/components/add-ssh-key-modal/add-ssh-key-modal.component';
import {SSHKeyEntity} from '../shared/entity/SSHKeyEntity';
import {UserGroupConfig} from '../shared/model/Config';

@Component({
  selector: 'kubermatic-sshkey',
  templateUrl: './sshkey.component.html',
  styleUrls: ['./sshkey.component.scss'],
})

export class SSHKeyComponent implements OnInit, OnDestroy {
  loading = true;
  sshKeys: SSHKeyEntity[] = [];
  sortedSshKeys: SSHKeyEntity[] = [];
  sort: Sort = {active: 'name', direction: 'asc'};
  userGroup: string;
  userGroupConfig: UserGroupConfig;
  private subscriptions: Subscription[] = [];
  projectID: string;

  constructor(
      private route: ActivatedRoute, private api: ApiService, private userService: UserService,
      private appConfigService: AppConfigService, public dialog: MatDialog) {}

  ngOnInit(): void {
    this.subscriptions.push(this.route.paramMap.subscribe((m) => {
      this.projectID = m.get('projectID');
      this.refreshSSHKeys();
    }));

    this.userGroupConfig = this.appConfigService.getUserGroupConfig();
    this.userService.currentUserGroup(this.projectID).subscribe((group) => {
      this.userGroup = group;
    });

    this.subscriptions.push(interval(5000).subscribe(() => {
      this.refreshSSHKeys();
    }));
    this.refreshSSHKeys();
  }

  ngOnDestroy(): void {
    for (const sub of this.subscriptions) {
      if (sub) {
        sub.unsubscribe();
      }
    }
  }

  refreshSSHKeys(): void {
    this.subscriptions.push(this.api.getSSHKeys(this.projectID).pipe(retry(3)).subscribe((res) => {
      this.sshKeys = res;
      this.sortSshKeyData(this.sort);
      this.loading = false;
    }));
  }

  addSshKey(): void {
    const dialogRef = this.dialog.open(AddSshKeyModalComponent);
    dialogRef.componentInstance.projectID = this.projectID;

    dialogRef.afterClosed().subscribe((result) => {
      result && this.refreshSSHKeys();  // tslint:disable-line
    });
  }

  trackSshKey(index: number, shhKey: SSHKeyEntity): number {
    const prevSSHKey = find(this.sshKeys, (item) => {
      return item.name === shhKey.name;
    });

    return prevSSHKey === shhKey ? index : undefined;
  }

  sortSshKeyData(sort: Sort): void {
    if (sort === null || !sort.active || sort.direction === '') {
      this.sortedSshKeys = this.sshKeys;
      return;
    }

    this.sort = sort;

    this.sortedSshKeys = this.sshKeys.sort((a, b) => {
      const isAsc = sort.direction === 'asc';
      switch (sort.active) {
        case 'name':
          return this.compare(a.name, b.name, isAsc);
        case 'status':
          return this.compare(a.spec.fingerprint, b.spec.fingerprint, isAsc);
        default:
          return 0;
      }
    });
  }

  compare(a, b, isAsc): number {
    return (a < b ? -1 : 1) * (isAsc ? 1 : -1);
  }
}
